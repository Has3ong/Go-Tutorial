package main

import (
	"fmt"
)

type size struct {
	length int "키"
	weight int "몸무게"
}

type People struct {
	name   string "이름"
	age    int    "나이"
	detail size   "신체사이즈"
}

func main() {
	people1 := People{
		"Kim",
		26,
		size{
			174,
			70,
		},
	}

	fmt.Println(people1)
}
