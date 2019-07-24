package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Printf("%p %v\n", &arr1, arr1)
	fmt.Printf("%p %v\n", &arr2, arr2)
}
