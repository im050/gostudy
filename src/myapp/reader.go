package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	var s = strings.NewReader("hello reader")
	var b = make([]byte, 8)
	for {
		i, err := s.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", i, err, b)
		fmt.Printf("b[:n] = %q\n", b[:i])

		if err == io.EOF {
			break;
		}
	}
}
