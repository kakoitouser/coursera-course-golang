/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
