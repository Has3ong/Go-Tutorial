package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


    j := 0
    for j < 5 {
        j++
    }

	arr := []string{"Alpha", "Bravo", "Charlie", "Delta"}
	for idx, char := range arr {
		fmt.Println(idx, char)
	}

	/*
	   for {
	       fmt.Println("Hello GoLang")
	   }
	*/
}
