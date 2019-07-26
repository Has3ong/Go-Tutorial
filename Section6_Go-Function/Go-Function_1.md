# Go-Function_1

#### Function example
```
func (function name) (parameter) (return type OR return name)

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}
```

* Multiple results
```
func swap(x, y string) (string, string) {
	return y, x
}
```

* Call Function
```
func sum(i int, f func(int, int)){
	f(i, 10)
}
func add(x int, y int){
	fmt.Println(x+y)
}
```

* Call by Reference, Call by value
```
func multi_value(x int) {
	x = x * 10
}

func multi_reference(x *int) {
	*x = *x * 10
}

func main() {
	sum(100, add)

	x := 100
	multi_value(x)
	fmt.Println(x)

	multi_reference(&x)
	fmt.Println(x)
}

```
