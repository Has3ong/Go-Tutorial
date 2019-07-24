package main

import "fmt"

func main() {
	var slice []int = make([]int, 5, 10)
	fmt.Println(slice, len(slice), cap(slice))

	var slice2 []int
	if slice2 == nil {
		fmt.Println("NIL")
	}

	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{}

	slice4 = slice3
	slice4[0] = 100
	fmt.Println(slice3)
	fmt.Println(slice4)

	for i, v := range slice3 {
		fmt.Println(i, v)
	}
}
