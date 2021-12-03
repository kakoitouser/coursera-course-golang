package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Uniq(input io.Reader, out io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		text := in.Text()
		if prev > text {
			return fmt.Errorf("data not sorted")
		}
		if prev == text {
			continue
		}
		fmt.Fprintln(out, text)
		prev = text
	}
	return nil
}

func main() {
	err := Uniq(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
	}
}
