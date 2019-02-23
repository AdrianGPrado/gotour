package main

import (
	"fmt"
	"time"
)

func printSomething(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("number ", i)
	}
}

func main() {
	go printSomething(10)
	printSomething(5)
}
