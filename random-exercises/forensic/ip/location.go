package ip

import (
	"github.com/oschwald/geoip2-golang"
	"net"
)

//todo - hide behind interface
type locationFetcher struct {
	db *geoip2.Reader
}

//todo - change filepath to relative after mods
func NewLocationFetcher() (*locationFetcher, error) {
	dbFileLocation := "/home/krzysiu/Downloads/GeoLite2-Country/GeoLite2-Country.mmdb"
	db, err := geoip2.Open(dbFileLocation)
	if err != nil {
		return nil, err
	}

	return &locationFetcher{db: db}, nil
}

func (f locationFetcher) Country(ip net.IP) (string, error) {
	country, err := f.db.Country(ip)
	if err != nil {
		return "", err
	}

	return country.Country.Names["en"], nil
}
