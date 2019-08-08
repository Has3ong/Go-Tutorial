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
	file, err := os.Open("TEST.txt")
	errCheck1(err)

	fileInfo, err := file.Stat()

	fd1 := make([]byte, fileInfo.Size())
	ct1, err := file.Read(fd1)

	fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name(), "\n")
	fmt.Println(fileInfo.Size(), "\n")
	fmt.Println(fileInfo.ModTime(), "\n")
	fmt.Println(ct1)
	fmt.Println(string(fd1))
}
