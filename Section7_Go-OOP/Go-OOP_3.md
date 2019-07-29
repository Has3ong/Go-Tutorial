# Go-OOP_3

#### Interface

```
/*
    type (interfacename) interface {

	}
*/
```

* If a structure is a collection of fields, an interface is a collection of methods.
* The interface defines the type of method prototype that must be implemented.

```
type Car struct {
	name   string
	speed int
}

type Behavior interface {
	Drive()
    Parking()
}

func (c Car) Drive() {
	fmt.Println(c.speed, " Driving")
}

func (c Car) Parking(){
    fmt.Println(c.name, " Parking")
}

func Print(i interface{}) {
	fmt.Println(i)
}
```

#### Check Interface Type

```
func main() {

	var a interface{}
	printType(a)

	a = 7.5
	printType(a)

	a = "Golang"
	printType(a)

	a = true
	printType(a)

	a = nil
	printType(a)
}

func printType(i interface{}) {
	fmt.Printf("%T\n", i)
}
```

#### Check Interface Type 2

```
type Car struct {
	name  string
	speed int
}

func Check(i interface{}) {
	switch i.(type) {
	case Car:
		object := i.(Car)
		fmt.Println("Car ", object.name, object.speed)

	case *Car:
		object := i.(*Car)
		fmt.Println("*Car ", object.name, object.speed)

	default:
		fmt.Println("Error")
	}
}
```