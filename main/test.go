package main

import "fmt" //alias
import "os"

var a string

func main() {

 	str := "Go says hello to the world!"

	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	fmt.Printf("Hello,world\n")

	fmt.Println("Hello,world")

	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")

	fmt.Println(&str);

	fmt.Println(HOME)
	fmt.Println(USER)
	fmt.Println(GOROOT)

	fmt.Printf("The operationg system user is: %s\n", USER)

	a = "G"
	print(a)
	f1()

}

func f1(){
	a := "O"
	print(a)
	f2()
}

func f2(){
	print(a)
}


