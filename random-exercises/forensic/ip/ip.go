package ip

import (
	"github.com/pkg/errors"
	"net"
	"regexp"
)

var ErrNoIPInStr = errors.New("string contains no ip address")

var Pattern = `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`

func Find(s string, r *regexp.Regexp) net.IP {
	ipStr := r.FindString(s)
	if ipStr == "" {
		return nil
	}
	return net.ParseIP(ipStr)
}

func QuantityOfCountries(ipSlice []net.IP, f locationFetcher) (map[string]int, error) {
	ret := make(map[string]int)

	for _, ip := range ipSlice {
		country, err := f.Country(ip)
		if err != nil {
			return nil, err
		}

		if _, ok := ret[country]; !ok {
			ret[country] = 1
			continue
		}

		ret[country]++
	}

	return ret, nil
}

//func QuantityOfDuplicatesMap(ipSlice []net.IP) map[string]int {
//	ret := make(map[string]int, len(ipSlice))
//
//	for _, ip := range ipSlice {
//		k := ip.String()
//
//		_, ok := ret[k]
//		if !ok {
//			ret[k] = 1
//			continue
//		}
//		ret[k]++
//	}
//
//	return ret
//}
