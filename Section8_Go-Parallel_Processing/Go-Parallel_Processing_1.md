# Go-Parallel_Processing_1

#### Go Routine

* Functions similar to threads in other languages
* Go-routine is a Lightweight logical thread managed by Go Runtime.
```
package main

import "fmt"
import "time"

func main() {

	fmt.Println("Main Start : ", time.Now())
	time.Sleep(4 * time.Second)
	fmt.Println("Main End : ", time.Now())
}
```

#### Multi-Core Utilization
* A goroutine is a lightweight thread managed by the Go runtime.
```
runtime.GOMAXPROCS(runtime.NumCPU())
//runtime.GOMAXPROCS(1)

for i := 0; i < 100; i++ {
	go exe(i) 
}
```