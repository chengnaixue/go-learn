package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v\n", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("GO")
}