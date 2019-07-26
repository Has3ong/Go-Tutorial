package main

import "fmt"

type People struct {
	name string
	age  int
}

func Age(p People) int {
	return p.age
}

func (p People) Age() int {
	return p.age + 10
}

type newInt int

func main() {
	Student := People{name: "A", age: 15}
	Doctor := People{name: "D", age: 45}

	fmt.Println(Student, Doctor)
	fmt.Println(Age(Student))
	fmt.Println(Student.Age())

	a := newInt(5)
	var b newInt = 15

	fmt.Println(a, b)

	var kim *People = new(People)
	kim.name = "kim"
	kim.age = 36

	fmt.Println(kim)

}
