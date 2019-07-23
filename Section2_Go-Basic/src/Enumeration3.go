package main

import "fmt"

func main() {

	const (
		_ = iota + 5*3
		A
		B
		C
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
