package main

import "fmt"

func main() {

	var p1 *int
	var p2 *int = new(int)

	fmt.Println(p1)
	fmt.Println(p2)

	a := 10
	p1 = &a
	p2 = &a

	fmt.Println(p1)
	fmt.Println(&p1)
	fmt.Println(*p1)

	fmt.Println(p2)
	fmt.Println(&p2)
	fmt.Println(*p2)

}
