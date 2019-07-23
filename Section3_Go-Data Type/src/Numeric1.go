package main

import "fmt"

func main() {
    var num1 int = 1
    var num2 int = -1
    var num3 int = 012  // 10
    var num4 int = 0x12 // 18
    fmt.Println(num1, num2, num3, num4)

	var ascii1 byte = 65
	var ascii2 byte = 0102
	var ascii3 byte = 0x43
	fmt.Printf("%c %c %c ", ascii1, ascii2, ascii3)

	var u1 rune = 44256
	var u2 rune = 0126340
	var u3 rune = 0xACE0
	fmt.Printf("%c %c %c ", u1, u2, u3)
}
