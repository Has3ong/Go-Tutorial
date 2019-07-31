package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", fmt.Errorf("%d", n)
}

func main() {

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
}
