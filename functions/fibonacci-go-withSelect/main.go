package main

import "fmt"

func fibonacci(ch chan int, quit chan string) {
	prev, current := 0, 1
	for {
		select {
		case ch <- prev:
			prev, current = current, prev+current
		case reason := <-quit:
			fmt.Println(reason)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
			if i == 9 {
				quit <- "exit"
			}
		}
	}()
	fibonacci(ch, quit)
}
