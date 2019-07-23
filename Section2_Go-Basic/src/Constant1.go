package main

import "fmt"

func main() {
	const _PI float32 = 3.14159265
	const _String = "Hello Go"

	/*
		const _Error string
		_Error = "Got Error"  // <- Error
	*/

	fmt.Println("PI : ", _PI)
	fmt.Println("String : ", _String)
}
