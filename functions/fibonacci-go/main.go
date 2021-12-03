package main

import "fmt"

func fibonacci(n int, ch chan int) {
	prev, current := 0, 1
	for i := 0; i < n; i++ {
		ch <- prev
		prev, current = current, prev+current
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	const n = 10

	go fibonacci(n, ch)
	for v := range ch {
		fmt.Println(v)
	}
}
