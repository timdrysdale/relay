package access

import (
	"regexp"

	"github.com/timdrysdale/relay/pkg/ttlcode"
)

type Options struct {
	disableAuth bool //not stable part of API

}

func NewOptions(cs *ttlcode.CodeStore) *Options {
	return &Options{}
}

// GetScopes returns a map of the scopes allowed for each
// path, with the path prefix as map key
// e.g. https://relay-access.yourdomain.io/session/8410928349108230498
// has the path prefix session

func getPrefixFromPath(path string) string {

	re := regexp.MustCompile("^\\/([\\w\\%-]*)\\/")

	matches := re.FindStringSubmatch(path)
	if len(matches) < 2 {
		return ""
	}

	// matches[0] = "/{prefix}/"
	// matches[1] = "{prefix}"
	return matches[1]

}
