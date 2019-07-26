# Go-OOP_2

#### Struct Initialize
```
type People struct {
	name string "이름"
	age  int    "나이"
}

people1 := struct{ name, job string }{"people1", "Student"}
peoples := []struct{ name, job string }{{"people1", "Doctor"}, {"people2", "Engineer"}}

people2 := reflect.TypeOf(People{})
```

#### Multiple Struct
```

type size struct {
	length int "키"
	weight int "몸무게"
}

type People struct {
	name   string "이름"
	age    int    "나이"
	detail size   "신체사이즈"
}

    people1 := People{
		"Kim",
		26,
		size{
			174,
			70,
		},
	}
```