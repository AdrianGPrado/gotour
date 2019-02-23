package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	right *Tree
// 	value int
// 	left  *Tree
// }

func main() {
	t1 := tree.New(2)
	fmt.Printf(t1)
}
