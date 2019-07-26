package main

import "fmt"

func stack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}()

	fmt.Println("Main End")

	stack()
}
