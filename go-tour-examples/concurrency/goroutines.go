package main

import (
	"fmt"
	"time"
)

func main() {
	say := func(s string) {
		for range [100]interface{}{} {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}

	// evaluation happens in the current goroutine and the execution happens in the new goroutine.
	go say("world")
	say("hello")
	
}
