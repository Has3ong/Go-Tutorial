package main

import "fmt"

func main() {
	var num1 complex64 = 2 + 1i
	num2 := complex64(2 + 1i)
	num3 := complex(2, 1)
	var num4 complex128 = 2 + 1i

	fmt.Println(num1, real(num1), imag(num1))
	fmt.Println(num2, real(num2), imag(num2))
	fmt.Println(num3, real(num3), imag(num3))
	fmt.Println(num4, real(num4), imag(num4))
}
