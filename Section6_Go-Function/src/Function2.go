package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 50)
}
func add(x int, y int) {
	fmt.Println(x + y)
}

func multi_value(x int) {
	x = x * 10
}

func multi_reference(x *int) {
	*x = *x * 10
}

func main() {
	sum(100, add)

	x := 100
	multi_value(x)
	fmt.Println(x)

	multi_reference(&x)
	fmt.Println(x)
}
