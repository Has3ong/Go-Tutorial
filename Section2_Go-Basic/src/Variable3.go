package main

import "fmt"

func main() {
	Test := 10
	Test2 := "Short Variables"

	//Test := 20 // <- Error
	//Test = 20 // <- Not Error
	fmt.Println(Test, Test2)

	if i := 0; i < 11 {
		fmt.Println("Short Variable Test : ", i)
	}
}
