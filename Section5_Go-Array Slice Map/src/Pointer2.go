package main

import "fmt"

func func1(n *int) {
	*n = 10
}

func func2(n int) {
	n = 100
}

func main() {

	a := 1

	func1(&a)
	fmt.Println(a)
	func2(a)
	fmt.Println(a)
	//func2(&a) // Error

}
