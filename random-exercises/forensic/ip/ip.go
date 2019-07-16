package ip

import (
	"github.com/pkg/errors"
	"net"
	"regexp"
)

var ErrNoIPInStr = errors.New("string contains no ip address")

//var ASDF = "hello"

var IPPattern = `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`

func Find(s string, r *regexp.Regexp) net.IP {
	ipStr := r.FindString(s)
	if ipStr == "" {
		return nil
	}
	return net.ParseIP(ipStr)
}
