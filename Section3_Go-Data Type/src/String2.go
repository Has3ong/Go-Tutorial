package main

import "fmt"

func main() {
	var str1 string = "Hello"
	for i, char := range str1 {
		fmt.Printf("%c %d\t", char, i)
	}
}
