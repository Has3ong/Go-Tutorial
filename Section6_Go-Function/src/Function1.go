package main

import "fmt"

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func main() {

	a, b := 10, 20
	c, d := "Hello", "Go-Lang"

	fmt.Println(add1(a, b))
	fmt.Println(add2(a, b))
	fmt.Println(swap(c, d))
}
