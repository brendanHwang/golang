// x,y int도 가능!
package main

import (
	"fmt"
)

func add(x, y int, z string) int {
	fmt.Println(z)
	return x + y
}

func main() {
	fmt.Println(add(42, 13, "test"))
}
