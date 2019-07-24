package main

import "fmt"

func main() {
	var str1 string = "Hello"

	fmt.Println(str1[0], str1[1], str1[2])            // 71 111 108
	fmt.Printf("%c %c %c", str1[0], str1[1], str1[2]) // H e l
}
