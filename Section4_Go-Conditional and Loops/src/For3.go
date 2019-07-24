package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			continue
		}
	}

	k := 0
	for {
		if k > 100 {
			break
		}
		k++
	}

Loop1:
	for l := 0; l < 5; l++ {
		for m := 0; m < 5; m++ {
			if l == 1 {
				break Loop1
			}
			fmt.Println(l, m)
		}
	}
}
