# Go-Parallel_Processing_3

#### Channel Advance

```
c := make(chan int)

go sendChannel(c, 10)          
go receiveChannel(c)    

func sendChannel(c chan<- int, cnt int) {
	for i := 0; i < cnt; i++ {
		c <- i
	}
    c <- 10000
	c <- 100
}

func receiveChannel(c <-chan int) {
	for i := range c {
		fmt.Println("received Channel Value : ", i)
	}
	fmt.Println(<-c)
}
```

#### Channel select case : 

```
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
```