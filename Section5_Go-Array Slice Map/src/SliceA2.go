package main

import "fmt"
import "sort"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1[0:9])
	fmt.Println(slice1[3:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[:])

	slice2 := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	slice3 := []string{"a", "d", "z", "ad", "bbacd", "dsdz"}

	sort.Ints(slice2)
	sort.Strings(slice3)

	fmt.Println(slice2)
	fmt.Println(slice3)

}
