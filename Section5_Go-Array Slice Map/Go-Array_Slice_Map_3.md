# Go-Array_Slice_Map_3

#### Slice Control
```
slice1 := []int{1, 2, 3, 4, 5}
slice1 = append(slice1, 6, 7, 8, 9, 10, 11, 12)
//slice1 [1 2 3 4 5 6 7 8 9 10 11 12]

slice2 := []int{1, 2, 3}
slice3 := []int{5, 6, 7}

slice4 := append(slice2, slice3...)
slice5 := append(slice2, slice3[0:1]...)

//slice4  [1 2 3 5 6 7]
//slice5  [1 2 3 5]

slice6 := make([]int, 0, 5)

for i := 0; i< 10; i++{
  slice6 = append(slice6, i)
}
```

#### Slice Extraction
```
slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice1[0:9] // [1 2 3 4 5 6 7 8 9]
slice1[3:]  // [4 5 6 7 8 9 10]
slice1[:5]  // [1 2 3 4 5]
slice1[:]   // [1 2 3 4 5 6 7 8 9 10]
```

#### Slice Sort
```
import "sort"

slice2 := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
slice3 := []string{"a", "d", "z", "ad", "bbacd", "dsdz"}

sort.Ints(slice2)
sort.Strings(slice3)
```

#### Slice Copy
* copy(destination, source)
* before the copy make()
```
slice1 := []int{1, 2, 3, 4, 5, 6}
slice2 := make([]int, 5)

copy(slice2, slice1)
slice2[0] = 100

slice1  // 1, 2, 3, 4, 5, 6
slice2  // 100, 2, 3, 4, 5
```
