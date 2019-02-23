package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond * 100)
	tack := time.After(time.Millisecond * 500)
	for {
		select {
		case <-tick:
			fmt.Printf(" tick ")
		case <-tack:
			fmt.Printf(" boom ")
		default:
			fmt.Printf("     .")
			time.Sleep(time.Millisecond * 50)
		}
	}
}
