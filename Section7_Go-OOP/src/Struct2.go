package main

import "fmt"

type newFunction func(int, int) int

func Output(x int, y int, fn newFunction) {
	fmt.Println(x, y, fn(x, y))
}
func main() {
	var test newFunction
	test = func(x int, y int) int {
		return (x * y)
	}

	Output(100, 250, test)
}
