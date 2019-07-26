package main

import "fmt"

func main() {
	count := increase()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}

func increase() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}
