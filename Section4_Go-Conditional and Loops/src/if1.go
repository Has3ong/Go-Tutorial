package main

import "fmt"

func main() {
	a := 10

	if a >= 10 {
		fmt.Println("True")
	}

	if a < 10 {
		fmt.Println("True")
	}

	//Error
	/*
	   if a < 10
	   {
	   }
	*/

	/*
	   if a < 10
	       fmt.Println()
	*/

	/*
	   if c := 1, c {
	       fmt.Println()
	   }
	*/
}
