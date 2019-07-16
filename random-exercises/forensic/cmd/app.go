package main

import (
	"fmt"
	"github.com/krzysztofSkolimowski/programming-practice/random-exercises/forensic/log"
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
	for i, v := range ips {
		fmt.Printf("%d.\t%v\n", i, v)
	}
}
