package main

import "fmt"

func main() {
	a := 10

	switch {
	case a > 8:
		fmt.Println("A")
	case a > 6 && a <= 8:
		fmt.Println("B")
	case a <= 6:
		fmt.Println("C")
	}
}
