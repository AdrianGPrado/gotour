package main

import "fmt"

func fibonacci(c, quit chan int) {
	a, b := 0, 1
	for {
		select {
		case c <- a:
			a, b = b, a+b
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
		for i := 0; i < 5; i++ {
			fmt.Printf("%v ", <-c)
		}
		quit <- 1
	}()
	fibonacci(c, quit)
}
