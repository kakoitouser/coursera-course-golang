package main

import "fmt"

func fibonacci() func() int {
	prev, current := 0, 1
	return func() int {
		res := prev
		prev, current = current, prev+current
		return res
	}
}
func main() {
	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
