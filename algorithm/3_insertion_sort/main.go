package main

import "fmt"

// 이미 앞에는 정렬되어 있다고 가정!
//삽입 정렬은 자기가 들어갈 위치를 찾는 과정
// 멀출 곳을 알고 있다!

// TODO insertion sort 구현

func insertion_sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		j := i
		for array[j] < array[j+1] {
			array[j], array[j+1] = array[j+1], array[j]
			j -= 1
		}
	}
}

func main() {
	array := []int{1, 10, 5, 8, 7, 6, 4, 3, 2, 9}
	InsertionSort(array)
	fmt.Print(array)
}
