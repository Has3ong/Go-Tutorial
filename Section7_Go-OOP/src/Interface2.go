package main

import "fmt"

func main() {

	var a interface{}
	printType(a)

	a = 7.5
	printType(a)

	a = "Golang"
	printType(a)

	a = true
	printType(a)

	a = nil
	printType(a)
}

func printType(i interface{}) {
	fmt.Printf("%T\n", i)
}
