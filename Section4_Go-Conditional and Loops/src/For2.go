package main

import "fmt"

func main() {
    for i, j := 0, 0; i i < 10; i, j = i+1, j + 5 {
        fmt.Println(i, j)
    }

    for k := 0; k < 5; k ++ {
        for l := 0; l < 5; l++ {
            fmt.Println(k, l)
        }
    }
}
