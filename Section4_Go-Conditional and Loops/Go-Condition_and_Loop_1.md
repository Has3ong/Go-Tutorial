# Go-Condition_and_Loop_1

#### If

```
if (condition) {
    (statement)
} else if (condition) {
    (statement)
} else {
    (statement)
}
```
```
a := 10

if a >= 10 {
    fmt.Println("True")
}

if a < 10 {
    fmt.Println("True")
}

Error
/*
    if a < 10
    {
    }
*/

/*
    if a < 10
        fmt.Println()
*/

/*
    if c := 1, c {
        fmt.Println()
    }
*/
 
```

#### if, else
```
a := 10

if a >= 10 {
    fmt.Println("True")
} else {
    fmt.Println("False")
}
```

#### if, else if, else
```
b := 100

if j >= 100 {
    fmt.Println("100 이상")
} else if j >= 50 && j < 100 {
    fmt.Println("50 이상 100 미만")
} else {
    fmt.Println("50 미만")
}
```