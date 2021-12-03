package main

import (
	"fmt"
	"strings"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func formatwork(in, i int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", goroutinesNum-in),
		"th", in,
		"iter", i, strings.Repeat("*", i))

}
func doSomeWork(in int) {
	for i := 0; i < iterationsNum; i++ {
		go fmt.Printf(formatwork(in, i))
	}
}
func main() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}
