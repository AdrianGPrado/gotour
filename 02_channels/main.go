package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}

func main() {
	s := []int{8, 4, 1, 9, -3, 1, 10, 3, 7}

	c := make(chan int)

	go sum(s[:4], c)
	go sum(s[4:], c)
	result := <-c + <-c
	fmt.Printf("Sum result: %v ", result)

}
