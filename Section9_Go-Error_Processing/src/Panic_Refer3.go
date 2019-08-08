package main

import "fmt"

func TestFunc() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}
	}()

	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 10; i++ {
		fmt.Println("index : ", arr[i])
	}
}

func main() {
	TestFunc()
	fmt.Println("Hello Golang!")
}
