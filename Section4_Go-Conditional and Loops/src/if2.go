package main

import "fmt"

func main() {
	a := 10
    if a >= 10 {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }

    b := 100
    if j >= 100 {
		fmt.Println("100 이상")
	} else if j >= 50 && j < 100 {
		fmt.Println("50 이상 100 미만")
	} else {
		fmt.Println("50 미만")
	}
}
