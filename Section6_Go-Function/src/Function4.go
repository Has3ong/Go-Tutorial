package main

import "fmt"

func multi(x, y int) (ret int) {
	ret = x * y
	return ret
}

func sum(x, y int) (ret int) {
	ret = x + y
	return ret
}

func main() {
	f1 := map[string]func(int, int) int{
		"multi_func": multi,
		"sum_func":   sum,
	}

	fmt.Println(f1["multi_func"](10, 10))
	fmt.Println(f1["sum_func"](15, 15))
}
