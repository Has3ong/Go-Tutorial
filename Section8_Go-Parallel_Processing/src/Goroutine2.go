package main

import "fmt"
import "time"
import "runtime"

func exe(name int) {

	fmt.Println(name, " Func start : ", time.Now())
	for i := 0; i < 100000; i++ {
	}
	fmt.Println(name, " Func end : ", time.Now())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	fmt.Println("Current Cpu : ", runtime.GOMAXPROCS(0))

	fmt.Println("Main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i)
	}
	fmt.Println("Main Routine End : ", time.Now())
}
