package main

import "fmt"
import "log"

func main() {
	fmt.Println("Start Main")
	log.Panic("Panic Occurred")
	fmt.Println("End Main")
}
