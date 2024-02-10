package main

import "fmt"

// https://www.acmicpc.net/problem/23882
// 1, 10 , 5, 8, 7, 6, 4, 3, 2, 9 를 오름차순으로 정렬하시오.

// NOTE] 슬라이스 자체는 포인터가 아니지만, 포인터를 사용하여 배열을 참조하는 방식으로 작동합니다.
// NOTE] O(N^2)의 시간 복작도

// TODO: 선택정렬 함수
func selectoin_sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		array[min], array[i] = array[i], array[min]
	}
}

// TODO: 입력 받기
func main() {
	array := []int{1, 10, 5, 8, 7, 6, 4, 3, 2, 9}
	selectoin_sort(array)
	fmt.Println(array)
}
