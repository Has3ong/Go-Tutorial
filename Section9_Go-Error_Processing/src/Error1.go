package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("unnamedfile")
	if err != nil {
		log.Fatal(err.Error())
		//log.Fatal(err)
	}
	fmt.Println(f.Name())
}
