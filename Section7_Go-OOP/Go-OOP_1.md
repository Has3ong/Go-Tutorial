# Go-OOP_1

#### User-Defined Type, Struct
* Go Lang has no concept of inheritance, class, or object.
```
type People struct {
    name string
    age int
}

// Method
func Age(p People) int {
	return p.age
}

// Struct People Method
func (p People) Age() int {
	return p.age + 10
}

var kim *People = new(People)
kim.name = "kim"
kim.age = 36

Student := People{name:"A", age : 15}
Doctor := People{name:"D", age: 45}

// {A 15} {D 45}
```

#### User-Defined Type, Data Type
```
type newInt int

a := cnt(5)
var b cnt = 15
```

#### User-Defined Type, Function
```
type newFunction func(int, int) int

func Output(x int, y int, fn newFunction) {
	fmt.Println(x, y, fn(x, y))
}

var test newFunction
	test = func(x int, y int) int {
		return (x * y)
	}

Output(100, 250, test)
// 100 250 25000
```

#### User-Defined Type Reference
```
type productList struct{ count, price int }

func (p productList) show() int {
	return p.count * p.price
}

func (p *productList) pointer_rePurchase(count, price int) {
	p.count += count
	p.price += price
}

func (p productList) rePurchase(count, price int) {
	p.count += count
	p.price += price
}
```