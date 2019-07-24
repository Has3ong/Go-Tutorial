package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := []int{}
	slice3 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice3[0][0] = 10
	/*
	   {10, 2, 3, 4},
	   {5, 6, 7, 8},
	*/
	fmt.Println(slice3)
}
