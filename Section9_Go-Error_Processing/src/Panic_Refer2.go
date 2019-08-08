package main

import "fmt"

func Test() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!")
	fmt.Println("test")

}

func main() {
	Test()
	fmt.Println("Hello Golang!")
}
