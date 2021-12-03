package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestUniq(t *testing.T) {
	testCases := []struct {
		need string
		in   string
	}{
		{"\r1\n2\n3\n4\n5\n", "\r1\n2\n3\n3\n3\n4\n5"},
	}

	for _, tc := range testCases {
		in := bufio.NewReader(strings.NewReader(tc.in))
		out := new(bytes.Buffer)
		err := Uniq(in, out)
		if err != nil {
			t.Errorf("test for Uniq func will be failed")
		}
		result := out.String()
		if result != tc.need {
			t.Errorf("test for Uniq fuunc will be failed \n we didnt get needed the desired result")
		}
		fmt.Println(result, tc.need)
	}
}
