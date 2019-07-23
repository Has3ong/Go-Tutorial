package main

import "fmt"

func main() {
	var num1 float32 = 3.14159267
	var num2 float32 = .1415926
	var num3 float32 = 1e-1
	var num4 float32 = 15e3
	var num5 float32 = 0.234E+3

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)

	var test float32 = 10.0
	fmt.Println(test - 0.1)
	fmt.Println(float32(test - (0.1)))
	fmt.Println(float64(test - (0.1)))
}
