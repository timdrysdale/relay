package shellrelay

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/timdrysdale/relay/pkg/access"
	"github.com/timdrysdale/relay/pkg/crossbar"
	"github.com/timdrysdale/relay/pkg/ttlcode"
)

func Relay(closed <-chan struct{}, parentwg *sync.WaitGroup, accessPort, relayPort int, audience, secret, target string, options access.Options) {

	var wg sync.WaitGroup

	var cs *ttlcode.CodeStore

	cs = ttlcode.NewDefaultCodeStore()

	config := crossbar.Config{
		Listen:    relayPort,
		Audience:  target,
		CodeStore: cs,
	}

	wg.Add(1)
	go crossbar.Crossbar(config, closed, &wg)

	wg.Add(1)
	go access.API(closed, &wg, accessPort, audience, secret, target, cs, options)

	wg.Wait()
	parentwg.Done()
	log.Trace("Relay done")
}
