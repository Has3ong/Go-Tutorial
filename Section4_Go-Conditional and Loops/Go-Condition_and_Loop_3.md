# Go-Condition_and_Loop_3

#### For
* The only loop that exists in the Go language
```
for (initilize);, (condition); (increse, decrese) {
    (statement)
}
```


```
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

j := 0
for j < 5 {
    j++
}
```

```
arr := []string{"Alpha", "Bravo", "Charlie", "Delta"}
for idx, char := range arr {
    fmt.Println(idx, char)
}
```

#### Infinite Loop
```
for {
    fmt.Println("Hello GoLang")
}
```

#### Nested loops
```
for i, j := 0, 0; i i < 10; i, j = i+1, j + 5 {
    fmt.Println(i, j)
}

for k := 0; k < 5; k ++ {
    for l := 0; l < 5; l++ {
        fmt.Println(k, l)
    }
}
```

#### Loop Control
```
for i := 0; i < 5; i ++ {
    if i % 2 == 0 {
        fmt.Println(i)
        continue
    }
}

k := 0
for {
    if k > 100{
        break
    }
    k ++
} 

Loop1:
    for l := 0; l < 5; l++ {
        for m := 0; m < 5; m++ {
            if l == 1{
                break Loop1
            }
        fmt.Println(l, m)
    }
}
```