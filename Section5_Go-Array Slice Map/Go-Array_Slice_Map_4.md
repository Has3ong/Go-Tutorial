# Go-Array_Slice_Map_4

#### Map
* Hash Table, Dictionary, Key-Value save a Data
* Reference Type
* Do not use Compare Operator
* Using Key is impossible but Value is possible
* No order
```
var map1 map[string] int = make(map[string]int)
var map2 = make(map[string]int)
map3 := make(map[string]int)
```

```
map1 := map[string]int{}
map1["Apple"] = 65
map1["Banana"] = 66
map1["Chocolate"] = 67

map2 := map[string]int{
  "Delta": 68,
  "Echo":  69,
  "Fox":   70,
}

map3 := make(map[string]int, 5)
map3["Golf"] = 71
map3["Hotel"] = 72
map3["Iota"] = 73
```

```
map1 := map[string]int{
  "Alpha": 65,
  "Bravo": 66,
  "Charlie": 67,
  "Delta": 68,
  "Echo":  69,
  "Fox":   70,
  "Golf": 71,
  "Hotel": 72,
}

// random output
for k, v := range map1{
  fmt.Println(k, v)
}
```

#### Map Control
```
map1 := map[string]int{
  "Go" : 1,
  "Python" : 2,
  "Java" : 3,
}
// Insert
map1["C#"] = 4

// Update
map1["C#"] = 5

// Delete
delete(map1, "C#")
```
