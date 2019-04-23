package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %v %v years \n", p.Name, p.Age)
}

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("address: %v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	a := Person{"Jan", 10}
	b := Person{"Krzyś", 15}
	c := Person{"Ania", 13}
	fmt.Println(a, b, c)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}


}
