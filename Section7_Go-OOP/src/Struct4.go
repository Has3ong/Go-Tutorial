package main

import (
	"fmt"
	"reflect"
)

type People struct {
	name string "이름"
	age  int    "나이"
}

func main() {

	people1 := struct{ name, job string }{"people1", "Student"}
	peoples := []struct{ name, job string }{{"people1", "Doctor"}, {"people2", "Engineer"}}

	fmt.Println(people1)
	fmt.Println(peoples)

	people2 := reflect.TypeOf(People{})
	fmt.Println(people2.Field(0), people2.Field(1))

}
