package main

import "fmt"

var number int = 9
var heap = [9]int{7, 6, 5, 8, 3, 5, 9, 1, 6}

func main() {
	for i := 1; i < number; i++ {
		for c := i; c != 0; {
			root := (c - 1) / 2
			if heap[root] < heap[c] {
				heap[root], heap[c] = heap[c], heap[root]
			}
			c = root
		}
	}
	fmt.Println(heap)

	for i := number - 1; i >= 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		root := 0
		c := 1
		for c < i {
			c = 2*root + 1
			if c < i-1 && heap[c] < heap[c+1] {
				c++
			}
			if c < i && heap[root] < heap[c] {
				heap[root], heap[c] = heap[c], heap[root]
			}

			root = c
		}
	}

	for _, v := range heap {
		fmt.Printf("%v ", v)
	}
}
