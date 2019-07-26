# Go-Function_3

#### Defer
```
func main() {
	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}()

	fmt.Println("Main End")
}
//Main End
//Defer 2
//Defer 1
```

#### Stack for Defer
```
func stack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
```

#### Closure
```
func main() {
	count := increase()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}

func increase() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}
```