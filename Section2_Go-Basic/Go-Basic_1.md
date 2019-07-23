# Go-Basic_1

#### Go Benefits and Features

* Simple grammar and simplicity
* Parallelism programming support
* Static type and dynamic execution
* Easy collaboration support
* Compile and run faster
* Does not handle generics and exceptions
* Convention unification

#### Declaring Variables, Constants

##### Variables Type
```
Variables type
intger, float, String, Boolean

etc
int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

##### Variables Declaration
```
var (Variables name) (Variables type)

var a int
var b, c, d int = 1, 2, 3
var e string = "Hello Golang"
var f = true
var g = 3.14

var (
    name string
    age int
    height float 32
    weight float32
    isWorking bool
)
```

##### Short Variables Declarations
* It is mainly used within a limited range of functions.
* An exception occurs when allocating after declaration.
```
(Variables name) := (Variables value)

Test := 10
Test2 := "Short Variables"

Test := 20 // <- Error 
Test = 20 // <- Not Error
```

##### Constants Declarations
* Once declared, the value can not be changed.
```
const (Constants name) (Constants type) = value

const _PI float32 = 3.14159265
const _String = "Hello Go"

const _Error string
_Error = "Got Error"  // <- Error


const (
    a, b int16 = 1, 2
    c, d = "Data", 1.23456
)
```
