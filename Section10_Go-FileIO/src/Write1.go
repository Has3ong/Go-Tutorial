package main

import (
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	file, err := os.Create("TEST.txt")
	errCheck1(err)
	string := "Hello Golang\n Hello VisualStudioCode"
	n2, err := file.WriteString(string)
	errCheck2(err)

	fmt.Println(n2)
	defer file.Close()
}
