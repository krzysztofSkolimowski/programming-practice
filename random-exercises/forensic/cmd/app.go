package main

import (
	"fmt"
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/ip"
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/log"
	"sort"
)

//func main() {
//	logFilePath := "/home/krzysiu/go/src/github.com/krzysztofSkolimowski" +
//		"/programming-practice/random-exercises/forensic/test-files/auth.log"
//
//	logs, err := log.FromFileFiltered(
//		logFilePath,
//		log.FailedLoginAttemptFilter,
//	)
//	if err != nil {
//		fmt.Println("Error: ", err)
//	}
//	for i, v := range logs {
//		fmt.Println("i: ", i, ". \t", v)
//	}
//}

func main() {
	logFilePath := "/home/krzysiu/go/src/github.com/krzysztofSkolimowski" +
		"/programming-practice/random-exercises/forensic/test-files/auth.log"

	ips, err := log.IPFromFileFiltered(
		logFilePath,
		log.FailedLoginAttemptFilter,
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	countryFetcher, err := ip.NewLocationFetcher()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	countriesQuantitiesMap, err := ip.QuantityOfCountries(ips, *countryFetcher)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	slice := CountryQuantitySlice{}
	for country, quantity := range countriesQuantitiesMap {
		slice = append(slice, CountryQuantityPair{country, quantity})
	}
	sort.Sort(slice)

	for i := 0; i < 10; i++ {
		v := slice[i]
		fmt.Println(i+1, ". \t|", v.quantity, "\t|", v.country)
	}
}

type CountryQuantitySlice []CountryQuantityPair

func (c CountryQuantitySlice) Len() int           { return len(c) }
func (c CountryQuantitySlice) Less(i, j int) bool { return c[i].quantity > c[j].quantity }
func (c CountryQuantitySlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

type CountryQuantityPair struct {
	country  string
	quantity int
}
