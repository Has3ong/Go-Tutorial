package main

import "fmt"

func main() {
	map1 := map[string]int{
		"Alpha":   65,
		"Bravo":   66,
		"Charlie": 67,
		"Delta":   68,
		"Echo":    69,
		"Fox":     70,
		"Golf":    71,
		"Hotel":   72,
	}

	// random output
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	map2 := map[string]int{
		"Go":     1,
		"Python": 2,
		"Java":   3,
	}
	// Insert
	map2["C#"] = 4
	fmt.Println(map2)

	// Update
	map2["C#"] = 5
	fmt.Println(map2)

	// Delete
	delete(map2, "C#")
	fmt.Println(map2)
}
