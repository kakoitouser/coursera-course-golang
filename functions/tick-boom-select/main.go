package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(120 * time.Millisecond)
	boom := time.Tick(600 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println(".")
			time.Sleep(60 * time.Millisecond)
		}
	}
}
