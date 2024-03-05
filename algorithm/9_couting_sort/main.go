package main

import "fmt"

func main() {
	var temp int
	count := [5]int{}
	array := [30]int{
		1, 3, 2, 4, 3, 2, 5, 3, 1, 2,
		3, 4, 4, 3, 5, 1, 2, 3, 5, 2,
		3, 1, 4, 3, 5, 1, 2, 1, 1, 1,
	}

	for i := 0; i < len(array); i++ {
		count[array[i]-1]++
	}
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			temp = i + 1
			fmt.Print(temp, " ")
		}
	}
}
