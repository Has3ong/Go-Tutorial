# Go-Function_2

#### Function example
```
func multiInt(n ...int) int {
	total := 1
	for _, value := range n {
		total *= value
	}

	return total
}

func multiString(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
}
```

#### Function call Slice
```
func multi(x, y int) (ret int) {
	ret = x * y
	return ret
}

func sum(x, y int) (ret int) {
	ret = x + y
	return ret
}

f := []func(int, int) int{multi, sum}
	num1 := f[0](12, 12)
	num2 := f[1](12, 12)
```

#### Function call Variables
```
func multi(x, y int) (ret int) {
	ret = x * y
	return ret
}

func sum(x, y int) (ret int) {
	ret = x + y
	return ret
}

var f1 func(int, int) int = multi
f2 := sum

fmt.Println(f1(10, 10), f2(10, 10))
```

#### Function with map
```
func multi(x, y int) (ret int) {
	ret = x * y
	return ret
}

func sum(x, y int) (ret int) {
	ret = x + y
	return ret
}

f1 := map[string]func(int, int) int{
	"multi_func": multi,
	"sum_func":   sum,
}

fmt.Println(f1["multi_func"](10, 10))
fmt.Println(f1["sum_func"](15, 15))
```

#### Recursive Function
* Factorial
```
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

x := fact(10)
```

#### Anonymous Function
```
func() {
    fmt.Println("Anoymous Function")
}()
```