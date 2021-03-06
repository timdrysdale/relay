package manifest

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/phayes/freeport"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	apiclient "github.com/timdrysdale/relay/pkg/bc/client"
	"github.com/timdrysdale/relay/pkg/bc/client/admin"
	groups "github.com/timdrysdale/relay/pkg/bc/client/groups"
	login "github.com/timdrysdale/relay/pkg/bc/client/login"
	"github.com/timdrysdale/relay/pkg/bc/models"
	"github.com/timdrysdale/relay/pkg/booking"
	"github.com/timdrysdale/relay/pkg/bookingstore"
	lit "github.com/timdrysdale/relay/pkg/login"
	"github.com/timdrysdale/relay/pkg/pool"
	"github.com/xtgo/uuid"
	"gopkg.in/yaml.v2"
)

var auth runtime.ClientAuthInfoWriter
var debug bool
var l *bookingstore.Limit
var ps *pool.PoolStore
var bookingDuration, mocktime, startime int64
var useLocal bool
var bearer, secret string
var bc *apiclient.Bc
var timeout time.Duration

func init() {

	useLocal = true

	debug = false

	if debug {
		os.Setenv("DEBUG", "true") //for apiclient
		log.SetReportCaller(true)
		log.SetLevel(log.TraceLevel)
		log.SetFormatter(&logrus.TextFormatter{FullTimestamp: false, DisableColors: true})
		defer log.SetOutput(os.Stdout)

	} else {
		os.Setenv("DEBUG", "false")
		log.SetLevel(log.WarnLevel)
		var ignore bytes.Buffer
		logignore := bufio.NewWriter(&ignore)
		log.SetOutput(logignore)
	}

}

func TestMain(m *testing.M) {

	var lt lit.Token
	var cfg *apiclient.TransportConfig
	var loginBearer string

	iat := time.Now().Unix() - 1
	nbf := iat
	exp := nbf + 30

	if useLocal {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		secret = "somesecret"

		ps = pool.NewPoolStore().
			WithSecret(secret).
			WithBookingTokenDuration(int64(180))

		l = bookingstore.New(ctx).WithFlush(time.Minute).WithMax(2).WithProvisionalPeriod(5 * time.Second)

		port, err := freeport.GetFreePort()
		if err != nil {
			panic(err)
		}

		host := "[::]:" + strconv.Itoa(port)

		audience := "http://" + host

		go booking.API(ctx, port, audience, secret, ps, l)

		time.Sleep(time.Second)

		lt = lit.NewToken(audience, []string{"everyone"}, []string{}, []string{"login:admin"}, iat, nbf, exp)
		loginBearer, err = lit.Signed(lt, secret)
		if err != nil {
			panic(err)
		}

		cfg = apiclient.DefaultTransportConfig().WithHost(host).WithSchemes([]string{"http"})

	} else { //remote

		remoteSecret, err := ioutil.ReadFile("../../secret/book.practable.io.secret")
		if err != nil {
			panic(err)
		}

		secret = strings.TrimSuffix(string(remoteSecret), "\n")

		lt = lit.NewToken("https://book.practable.io", []string{"everyone"}, []string{}, []string{"login:admin"}, iat, nbf, exp)
		loginBearer, err = lit.Signed(lt, secret)
		if err != nil {
			panic(err)
		}

		cfg = apiclient.DefaultTransportConfig().WithSchemes([]string{"https"})

	}

	loginAuth := httptransport.APIKeyAuth("Authorization", "header", loginBearer)

	bc = apiclient.NewHTTPClientWithConfig(nil, cfg)

	timeout = 10 * time.Second

	params := login.NewLoginParams().WithTimeout(timeout)
	resp, err := bc.Login.Login(params, loginAuth)
	if err != nil {
		panic(err)
	}

	bearer = *resp.GetPayload().Token

	auth = httptransport.APIKeyAuth("Authorization", "header", bearer)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestExample(t *testing.T) {
	exp := int64(1613256113)
	m0 := Example(exp)

	_, err := yaml.Marshal(m0)

	assert.NoError(t, err)

	content, err := ioutil.ReadFile("testdata/example.yaml")
	assert.NoError(t, err)

	m1 := &Manifest{}
	err = yaml.Unmarshal(content, m1)
	assert.NoError(t, err)

	assert.Equal(t, m0, m1)

}

// TestProduceExample is only used to generate the example
// after a breaking update, and should be checked by hand
// This "fails" if enabled
func testProduceExample(t *testing.T) {
	// This fails if enabled"

	exp := int64(1613256113)
	m0 := Example(exp)

	buf, err := yaml.Marshal(m0)

	assert.NoError(t, err)

	err = ioutil.WriteFile("testdata/example.yaml", buf, 0644)

	assert.NoError(t, err)

	if err == nil {
		fmt.Println("Examples file successfully written to testdata/example.yaml")
	}

	// ensure we notice if enabled accidentally
	t.Fatal("wrote example file")

}

func TestBearer(t *testing.T) {
	// admin
	token, err := jwt.ParseWithClaims(bearer, &lit.Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	assert.NoError(t, err)

	claims, ok := token.Claims.(*lit.Token)

	assert.True(t, ok)
	assert.True(t, token.Valid)

	assert.Equal(t, []string{"everyone"}, claims.Groups)
	assert.Equal(t, []string{"booking:admin"}, claims.Scopes)
	assert.True(t, claims.ExpiresAt > time.Now().Unix()+30)

	_, err = uuid.Parse(claims.Subject)
	assert.NoError(t, err)

}

func TestAddGroup(t *testing.T) {

	n := "somegroup"
	ty := "group"
	g := &models.Group{
		Description: &models.Description{
			Name: &n,
			Type: &ty,
		},
	}
	params := groups.NewAddNewGroupParams().
		WithTimeout(timeout).
		WithGroup(g)

	resp, err := bc.Groups.AddNewGroup(params, auth)

	assert.NoError(t, err)

	gid := resp.GetPayload()

	if debug {
		fmt.Println(*gid.ID)
	}

	_, err = uuid.Parse(*gid.ID)

	assert.NoError(t, err, "groupID not a uuid")

}

func TestUploadManifest(t *testing.T) {

	if useLocal != true {
		t.Fatal("Don't run this test on a live server - it wipes the everything!")
	}

	// clear store
	err := bc.Admin.DeletePoolStore(
		admin.NewDeletePoolStoreParams().
			WithTimeout(timeout),
		auth)

	assert.Error(t, err)

	assert.Equal(t, "[DELETE /admin/poolstore][404] deletePoolStoreNotFound  <nil>", err.Error())

	m := Example(time.Now().Unix() + 3600)

	status, err := UploadManifest(bc, auth, timeout, *m)

	assert.NoError(t, err)

	assert.Equal(t, int64(2), status.Groups)
	assert.Equal(t, int64(3), status.Pools)
	assert.Equal(t, int64(6), status.Activities)

	if debug {
		pretty, err := json.MarshalIndent(status, "", "\t")
		assert.NoError(t, err)
		fmt.Println(string(pretty))
	}

}
