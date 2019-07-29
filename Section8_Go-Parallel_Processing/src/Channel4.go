package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- 1
		}
		close(ch)
	}()

	val1, ch1 := <-ch
	fmt.Println(val1, ch1)
	val2, ch2 := <-ch
	fmt.Println(val2, ch2)
	val3, ch3 := <-ch
	fmt.Println(val3, ch3)
	val4, ch4 := <-ch
	fmt.Println(val4, ch4)
}
