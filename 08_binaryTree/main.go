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
	if t.Right != nil {
		walk(t.Right, ch)
	}
	ch <- t.Value
}

func walker(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func printTree(ch chan int) {
	for i := range ch {
		fmt.Printf("%v ", i)
	}
}

func compareTrees(ch1, ch2 chan int) bool {
	for i := range ch1 {
		j := <-ch2
		if i != j {
			return false
		}
	}
	return true
}

func main() {
	ch1 := make(chan int)
	t1 := tree.New(2)
	go walker(t1, ch1)
	t2 := tree.New(2)
	ch2 := make(chan int)
	go walker(t2, ch2)

	fmt.Println("\ntree 1 ")
	printTree(ch1)

	fmt.Println("\ntree 2 ")
	printTree(ch2)

	equal := compareTrees(ch1, ch2)
	fmt.Println("\nAre trees equal: ", equal)

}
