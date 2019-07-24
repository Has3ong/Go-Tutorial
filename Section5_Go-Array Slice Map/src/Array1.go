package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4} // 1, 2, 3, 4, 0

	// Array size auto initialize
	arr5 := [...]int{1, 2, 3, 4, 5}
	arr6 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	arr7 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(len(arr7), cap(arr7))
}
