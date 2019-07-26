package main

import "fmt"

func multiInt(n ...int) int {
	total := 1
	for _, value := range n {
		total *= value
	}

	return total
}

func multiString(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
}

func multi(x, y int) (ret int) {
	ret = x * y
	return ret
}

func sum(x, y int) (ret int) {
	ret = x + y
	return ret
}

func main() {
	x := multiInt(5, 6, 7, 8, 9, 10)
	fmt.Println(x)

	multiString("Hello", "World", "Golang", "Github")

	f := []func(int, int) int{multi, sum}
	num1 := f[0](12, 12)
	num2 := f[1](12, 12)

	fmt.Println(num1, num2)

	var f1 func(int, int) int = multi
	f2 := sum

	fmt.Println(f1(10, 10), f2(10, 10))
}
