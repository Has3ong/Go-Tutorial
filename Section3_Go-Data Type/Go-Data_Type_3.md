# Go-Data_Type_2

#### String Type
* Character Type -> rune(int32)

```
var str string = "Hello Golang"
var unicode8 string = "\uace0"
```

#### String
```
var str1 string = "Hello"
var str2 string = "Golang"

fmt.Println(str1[0], str1[1], str1[2]) // 71 111 108
fmt.Printf("%c %c %c", str1[0], str1[1], str1[2]) // H e l
```

#### Escape
```

```

#### with for loop
```
var str1 string = "Hello"
for i, char := range str1 {
  fmt.Printf("%c %d\t", char, i)
}
```

#### String Slicing
```
var str string = "Hello GoLang"

str[0:3]    // Hel
str[3:]     // lo GoLang
```

#### with Compare Operator
```
var str1 string = "Hello"
var str2 string = "GoLang"

str1 == str2    // false
str1 != str2    // true 
str1 > str2     // true
str1 < str2     // false

// Comapre with ASCII Code
```

#### String Control
```
str1 := "Go is an open source programming language that makes it easy to build" 
str2 := "simple, reliable, and efficient software."

str1 + str2 // Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

strAppend := []string{}
strAppend = append(strAppend, str1)
strAppend = append(strAppend, str2)

strAppend // [Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.]
```

