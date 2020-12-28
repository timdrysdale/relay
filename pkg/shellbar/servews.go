package shellbar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eclesh/welford"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/timdrysdale/relay/pkg/permission"
	"github.com/timdrysdale/relay/pkg/util"
)

type ConnectionType int

const (
	Session ConnectionType = iota
	Shell
	Unsupported
)

// serveWs handles websocket requests from clients.
func serveWs(closed <-chan struct{}, hub *Hub, w http.ResponseWriter, r *http.Request, config Config) {

	// check if topic is of a supported type before we go any further
	ct := Unsupported

	path := slashify(r.URL.Path)

	log.WithField("path", path).Trace()

	connectionType := getConnectionTypeFromPath(path)
	topic := getTopicFromPath(path)

	log.Trace(fmt.Sprintf("%s -> %s and %s\n", path, connectionType, topic))

	if connectionType == "shell" {
		ct = Shell
	}

	log.WithField("connectionType", ct).Trace()

	if ct == Unsupported {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.WithField("connectionType", connectionType).Error("connectionType unsuported")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithField("error", err).Error("serveWs failed to upgrade to websocket")
		return
	}

	log.Trace("upgraded to ws") //Cannot return any http responses from here on

	// Enforce permissions by exchanging the authcode for a connection ticket
	// which contains expiry time, route, and permissions

	// Get the first code query param, lowercase only
	var code string

	code = r.URL.Query().Get("code")

	log.WithField("code", code).Trace()

	// if no code or empty, return 401
	if code == "" {
		log.WithField("topic", topic).Info("Unauthorized - No Code")
		return
	}

	// Exchange code for token

	token, err := config.CodeStore.ExchangeCode(code)

	log.WithFields(log.Fields{"token": util.Compact(token), "error": err}).Trace("Exchange code")

	if err != nil {
		log.WithField("topic", topic).Info("Unauthorized - Invalid Code")
		return
	}

	// check token is a permission token so we can process it properly
	// It's been validated so we don't need to re-do that
	if !permission.HasRequiredClaims(token) {
		log.WithField("topic", topic).Info("Unauthorized - original token missing claims")
		return
	}

	now := config.CodeStore.GetTime()

	if token.NotBefore > now {
		log.WithField("topic", topic).Info("Unauthorized - Too early")
		return
	}

	ttl := token.ExpiresAt - now

	log.WithFields(log.Fields{"ttl": ttl, "topic": topic}).Trace()

	audienceBad := (config.Audience != token.Audience)
	topicBad := (topic != token.Topic)
	expired := ttl < 0

	if audienceBad || topicBad || expired {
		log.WithFields(log.Fields{"audienceBad": audienceBad, "topicBad": topicBad, "expired": expired, "topic": topic}).Trace("Token invalid")
		return
	}

	// check permissions

	var canRead, canWrite bool

	for _, scope := range token.Scopes {
		if scope == "read" {
			canRead = true
		}
		if scope == "write" {
			canWrite = true
		}
	}

	if !(canRead || canWrite) {
		log.WithFields(log.Fields{"topic": topic, "scopes": token.Scopes}).Trace("No valid scopes")
		return
	}

	cancelled := make(chan struct{})

	// cancel the connection when the token has expired
	go func() {
		time.Sleep(time.Duration(ttl) * time.Second)
		close(cancelled)
	}()

	if ct == Shell {
		// initialise statistics
		tx := &Frames{size: welford.New(), ns: welford.New()}
		rx := &Frames{size: welford.New(), ns: welford.New()}
		stats := &Stats{connectedAt: time.Now(), tx: tx, rx: rx}

		client := &Client{hub: hub,
			conn:       conn,
			send:       make(chan message, 256),
			topic:      topic + token.TopicSalt,
			stats:      stats,
			name:       uuid.New().String(),
			userAgent:  r.UserAgent(),
			remoteAddr: r.Header.Get("X-Forwarded-For"),
			audience:   config.Audience,
			canRead:    canRead,
			canWrite:   canWrite,
		}
		client.hub.register <- client

		log.WithField("Topic", client.topic).Trace("Registering Client")

		go client.writePump(closed, cancelled)
		go client.readPump()

		if token.AlertHost {
			// initialise statistics
			tx := &Frames{size: welford.New(), ns: welford.New()}
			rx := &Frames{size: welford.New(), ns: welford.New()}
			stats := &Stats{connectedAt: time.Now(), tx: tx, rx: rx}

			// alert SSH host agent to make a new connection to relay at the same address
			adminClient := &Client{
				topic: getHostTopicFromUniqueTopic(topic),
				name:  uuid.New().String(),
				stats: stats,
			}

			hub.register <- adminClient

			permission.SetAlertHost(&token, false) //turn off host alert
			code = config.CodeStore.SubmitToken(token)
			log.WithField("token", token).Trace("Submitting token for host")

			// same URL as client used, but different code (and leave out the salt)
			hostAlertURI := token.Audience + "/" + token.ConnectionType + "/" + token.Topic + "?code=" + code
			client.hostAlertURI = hostAlertURI
			ca := ConnectionAction{
				Action: "connect",
				URI:    hostAlertURI,
			}

			camsg, err := json.Marshal(ca)

			if err != nil {
				log.WithFields(log.Fields{"error": err, "uri": hostAlertURI}).Error("Failed to make connectionAction message")
				return
			}

			time.Sleep(100 * time.Millisecond)
			hub.broadcast <- message{sender: *adminClient, data: camsg, mt: websocket.TextMessage}
			time.Sleep(100 * time.Millisecond)
			hub.unregister <- adminClient

		}

		return
	}

}
