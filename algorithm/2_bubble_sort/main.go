package main

import "fmt"

//NOTE] 오름차순 버블정렬 한번 반복 할 때 가장 큰 값을 뒤로 보낸다!

// TODO 버블정렬 구현
func bubble_sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1+-i; j++ {
			if array[j+1] < array[j] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// TODO 입력

func main() {
	array := []int{1, 10, 5, 8, 7, 6, 4, 3, 2, 9}
	bubble_sort(array)
	fmt.Print(array)
}
