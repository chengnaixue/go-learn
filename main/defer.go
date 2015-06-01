package main

import "fmt"

func main() {
	Function1()
	Function3()
}

func Function1() {
	fmt.Printf("In Function1 at the top\n")
	defer Function2()
	fmt.Printf("In Function2 at the bottom!\n")
}

func Function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

func Function3() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
