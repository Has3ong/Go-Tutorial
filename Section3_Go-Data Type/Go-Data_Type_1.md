# Go-Data_Type_1

#### Boolean Type
* True and False
* Primarily used with conditional logical operators. !, || (or), && (and)
* No implicit casting.
```
var bool1 bool = true
var bool2 bool = false
bool3 := true
bool4 := false
```

```
bool1 == bool3
bool1 && bool2
bool1 || bool2
!bool 3
```

```
var bool5 int = 1
if bool5 {
    fmt.Println(bool5)
}

// Error Occured
// non-bool bool5 (type int) used as if condition
```

##### with Logical Operators
```
true && true : true
true && false : false
false && false : false

true || ture : true
true || false : true
false || false : false

!true : false
!false : true
```

##### with Compare Operators
```
num1, num2 := 1, 2

num1 < num2 : true
num1 <= num2 : false
num1 > num2 : false
num1 >= num2 : false
num1 == num2 : false
num1 != num2 : true
```