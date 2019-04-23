package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fib() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, second+first
		return second
	}
}

func main() {
	fmt.Println("===============================================")
	pos, neg := adder(), adder()

	for _, i := range []int{1, 1, 1, 1, 1} {
		fmt.Println(pos(i), neg(-2*i))
	}
	fmt.Println("===============================================")

	fib := fib()

	for range [20]interface{}{} {
		fmt.Println(fib())
	}

	fmt.Println("===============================================")
}
