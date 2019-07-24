# Go-Array_Slice_Map_1

#### Array

```
var arr1 [5]int // {0, 0, 0, 0, 0}
var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
var arr3 = [5]int{1, 2, 3, 4, 5}
arr4 := [5]int{1, 2, 3, 4} // 1, 2, 3, 4, 0

// Array size auto initialize
arr5 := [...]int{1, 2, 3, 4, 5}
arr6 := [2][2]int{
  {1, 2},
  {3, 4},
}
```

* len(), cap()
```
arr7 := [5]int{1, 2, 3, 4, 5}
len(arr7), cap(arr7)  // 5, 5
```

#### Array Search
```
arr1 := [5]int{1, 2, 3, 4, 5}

for i := 0; i < len(arr1); i++ {
  fmt.Println(arr1[i])
}

for j, v := range arr1 {
  fmt.Println(j, v)
}

/*
  for _, v := range arr1 {
    fmt.Println(v) // print value
  }

  for v, j := range arr1 {
    fmt.Println(v) // print index
  }
*/
```

#### Array Copy
```
arr1 := [5]int{1, 2, 3, 4, 5}
arr2 := arr1

// arr2 = {1, 2, 3, 4, 5}
```
