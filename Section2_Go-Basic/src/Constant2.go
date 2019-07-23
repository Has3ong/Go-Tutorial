package main

import "fmt"

func main() {
	const (
		a, b int16 = 1, 2
		c, d       = "Data", 1.23456
	)

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "d : ", d)
}
