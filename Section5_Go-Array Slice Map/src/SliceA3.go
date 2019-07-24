package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := make([]int, 5)

	copy(slice2, slice1)
	slice2[0] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)
}
