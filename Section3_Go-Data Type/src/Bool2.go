package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

    num1, num2 := 1, 2

	fmt.Println(num1 < num2)
	fmt.Println(num1 <= num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
}
