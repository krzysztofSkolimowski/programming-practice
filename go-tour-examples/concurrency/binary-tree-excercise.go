package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		walk(t.Left, ch)
	}

	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}

}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	myTree := tree.New(1)
	fmt.Println(myTree)
	ch := make(chan int)

	go Walk(myTree, ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(myTree, tree.New(1)))
	fmt.Println(Same(myTree, tree.New(2)))
	fmt.Println(Same(myTree, tree.New(3)))
}
