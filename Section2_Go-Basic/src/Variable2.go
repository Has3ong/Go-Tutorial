package main

import "fmt"

func main() {
	var (
		name      string
		age       int
		height    float32
		weight    float32
		isWorking bool
	)

	name = "Has3ong"
	age = 20
	height = 180.34
	weight = 86.56
	isWorking = true

	fmt.Println("name : ", name, "age : ", age, "height : ", height, "weight : ", weight, "isWorking : ", isWorking)
}
