package main

import (
	"fmt"
	"sync"
	"time"
)

type num struct {
	mu     sync.Mutex
	number int
}

func (n *num) Increment() {
	n.mu.Lock()
	n.number++
	n.mu.Unlock()
}

func (n *num) getNum() int {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.number
}

func main() {
	n := num{number: 0}
	for i := 0; i < 10; i++ {
		go n.Increment()
		fmt.Println(n.getNum())
	}
	time.Sleep(2 * time.Second)
}
