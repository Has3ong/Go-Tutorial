package main

import "fmt"

func receiveChannel(cnt int) <-chan int {
	sum := 0
	channel := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		channel <- sum
		channel <- 1000
		channel <- 200000
		close(channel)
	}()
	return channel
}

func TotalSum(c <-chan int) <-chan int {
	total := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		total <- a
	}()
	return total
}

func main() {

	channel1 := receiveChannel(10)
	output := TotalSum(channel1)

	fmt.Println(<-output)
}
