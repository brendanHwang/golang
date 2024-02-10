package main

import (
	"fmt"
	"freeCodeCamp/ch3/page12"
)

func main() {
	//fmt.Println(page5.GetCoord())
	divide, err := page12.Divide(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(divide)
}
