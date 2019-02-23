package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	val map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.val[key]++
}
func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.val[key]
}
func main() {
	sf := SafeCounter{val: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sf.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println("value: ", sf.Value("somekey"))
}
