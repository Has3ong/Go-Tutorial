package main

import "fmt"
import "time"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 100
			time.Sleep(1)
		}
	}()

	go func() {
		for {
			ch2 <- "Go-lang Channel !"
			time.Sleep(1)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("case num : ", num)
			case str := <-ch2:
				fmt.Println("case str : ", str)
				//default:
				//fmt.Println("default test")
			}
		}
	}()

	time.Sleep(3 * time.Second)
}
