package main

import (
	"fmt"
	"time"
)

func sendChannel(c chan<- int, cnt int) <-chan int {
	for i := 0; i < cnt; i++ {
		c <- i
	}
	c <- 10000
	c <- 100
}

func receiveChannel(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)

	go sendChannel(c, 10)
	go receiveChannel(c)
	time.Sleep(2 * time.Second)
}
