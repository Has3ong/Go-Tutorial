package main

import (
	"fmt"
)

func Sum(rg int, c chan int) {
	sum := 0

	for i := 1; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

func main() {

	c := make(chan int)

	go Sum(10000, c)
	go Sum(3000, c)
	go Sum(5000, c)

	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
