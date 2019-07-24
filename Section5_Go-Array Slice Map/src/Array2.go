package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for j, v := range arr1 {
		fmt.Println(j, v)
	}

	/*
	  for _, v := range arr1 {
	    fmt.Println(v) // print value
	  }

	  for v, j := range arr1 {
	    fmt.Println(v) // print index
	  }
	*/
}
