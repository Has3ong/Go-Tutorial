package main

import "fmt"

func main() {

	const (
		A = iota + 1
		B
		C
		_
		D
		E
		F
	)

	fmt.Println("A : ", A)
	fmt.Println("B : ", B)
	fmt.Println("C : ", C)
	fmt.Println("D : ", D)
	fmt.Println("E : ", E)
	fmt.Println("F : ", F)
}
