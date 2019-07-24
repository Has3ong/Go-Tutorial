package main

import "fmt"

func main() {

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

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
}
