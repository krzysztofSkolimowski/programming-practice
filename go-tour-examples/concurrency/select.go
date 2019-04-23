package main

import (
	"fmt"
	"time"
)

func fibonacciQuit(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 30; i++ {
		//for {
			fmt.Println(<-c)
			time.Sleep(time.Second/2)
		}
		quit <- 0
	}()
	fibonacciQuit(c, quit)

}
