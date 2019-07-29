# Go-Parallel_Processing_2

#### Channel

* Go channel is a channel that sends and receives data through the channel.
* Used for data exchange and execution flow synchronization between Go Routine.

```
func work1(v chan int) {
	fmt.Println("Work1 : Start : ", time.Now())
	fmt.Println("Work1 : End :  ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : Start : ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : End : ", time.Now())
	v <- 2
}

v := make(chan int)
go work1(v)
go work2(v)

<-v
<-v
```

#### Channel

```
func Sum(rg int, c chan int) {
	sum := 0

	for i := 1; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

c := make(chan int)

go Sum(10000, c)
go Sum(3000, c)
go Sum(5000, c)

result1 := <-c
result2 := <-c
result3 := <-c

fmt.Println(result1)
fmt.Println(result2)
fmt.Println(result3)
```

#### Using Channel Buffer

```
//Using Buffer
//ch := make(chan bool, 4)
ch := make(chan bool)
```

#### Channel Close

```
ch := make(chan int)
go func() {
	for i := 0; i < 3; i++ {
		ch <- 1
	}
	close(ch)
}()
```