# Go-Array_Slice_Map_2

#### Slice
```
var slice1 []int
slice2 := []int{}
slice3 := [][]int{
  {1, 2, 3, 4},
  {5, 6, 7, 8},
}

slice3[0][0] = 10
/*
  {10, 2, 3, 4},
  {5, 6, 7, 8},
*/
```

#### Using make
```
// slice's length is 5, physical memory size is 10
var slice []int = make([]int, 5, 10)

//nil -> length, size is 0
var slice2 []int
if slice2 == nil{
  fmt.Println("NIL")
}
```

#### Slice Reference
```
slice3 := []int{1, 2, 3, 4, 5}
slice4 := []int{}

slice4 = slice3
slice4[0] = 100
```

#### Slice Search
```
for i, v := range slice3{
  fmt.Println(i, v)
}
```
