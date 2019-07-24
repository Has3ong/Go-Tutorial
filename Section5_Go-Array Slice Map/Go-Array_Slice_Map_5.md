# Go-Array_Slice_Map_5

#### Pointer
* A pointer holds the memory address of a value.
* The type *T is a pointer to a T value. Its zero value is nil.

```
var p1 *int             // <nil>
var p2 *int = new(int)  // 0xc000092000

a := 10
p1 = &a // address value
p2 = &a

p1 &a // 0xc00009e000
&p1   // 0xc00000e020
*p1   // 10

p2 &a // 0xc00009e000
&p2   // 0xc00000e018
*p2   // 10
```

#### Pointer Advance, Function
```
func func1(n *int){
  *n = 100
}

func func2(n int){
  n = 100
}

a := 10

func1(&a)
func2(a)
//func2(&a) // Error
```
