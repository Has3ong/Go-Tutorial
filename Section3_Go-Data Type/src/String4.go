package main

import (
	"fmt"
)

func main() {
	str1 := "Go is an open source programming language that makes it easy to build"
	str2 := "simple, reliable, and efficient software."

	strAppend := []string{}
	strAppend = append(strAppend, str1)
	strAppend = append(strAppend, str2)

	fmt.Println(strAppend)
}
