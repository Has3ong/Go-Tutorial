## Go-Package_Basic_1

#### Init method
```
func init() {
	fmt.Println("Init Start")
}

func main() {
	fmt.Println("Main Start")
}
```

```
package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init Method 1 ")
}

func init() {
	fmt.Println("Init Method 2")
}

func init() {
	fmt.Println("Init Method 3")
}

func main() {
	fmt.Println("Main Method")
}

```