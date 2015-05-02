package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	var intP *int = &i1
	var intP2 *int
	fmt.Printf("An integer: %d, it's location in memory: %p \n", i1, intP)
	fmt.Printf("%d\n", *intP)
	fmt.Printf("%p\n", intP2)
}
