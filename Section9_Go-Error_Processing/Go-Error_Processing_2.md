## Panic Refer_2

#### Panic

* The function Panic() as the function Go built-in stops the current function immediately and returns immediately after all the drop functions have been performed on the current function.
* This panic mode practice again applies equally to higher functions and continues to climb up the colstack.
* And at the end, the program ends with an error.

#### Recover

* The function Recover() as a function of Go internal function is a function of returning the panic state by the panic function to normal state.
* In the above example of panic, the main function fails to call the println() and the program crash, but using the cover function, as shown below and in the example, removes the panic state and invokes the next sentence of openFile().
```
func Test() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!")
	fmt.Println("test")

}

func main() {
	Test()
	fmt.Println("Hello Golang!")
}
```

```
func TestFunc() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}
	}()

	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 10; i++ {
		fmt.Println("index : ", arr[i])
	}
}

func main() {
	TestFunc()
	fmt.Println("Hello Golang!")
}
```