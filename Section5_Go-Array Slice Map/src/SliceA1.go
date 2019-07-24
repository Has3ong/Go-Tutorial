package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice1 = append(slice1, 6, 7, 8, 9, 10, 11, 12)

	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := []int{1, 2, 3}
	slice3 := []int{5, 6, 7}

	slice4 := append(slice2, slice3...)
	slice5 := append(slice2, slice3[0:1]...)

	fmt.Println(slice4)
	fmt.Println(slice5)

	slice6 := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		slice6 = append(slice6, i)
		fmt.Println(len(slice6), cap(slice6), slice6)
	}

}
