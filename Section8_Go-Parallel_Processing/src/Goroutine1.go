package main

import "fmt"
import "time"

func main() {
	fmt.Println("Main Start : ", time.Now())
	time.Sleep(4 * time.Second)
	fmt.Println("Main End : ", time.Now())
}
