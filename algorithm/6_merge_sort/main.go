package main

import "fmt"

var number = 8
var sorted = [8]int{}

func merge(a []int, m int, middle int, n int) {
	i := m
	j := middle + 1
	k := m

	// 작은 순서대로 배열에 삽입
	for i <= middle && j <= n {
		if a[i] <= a[j] {
			sorted[k] = a[i]
			i++
		} else {
			sorted[k] = a[j]
			j++
		}
		k++
	}

	for i <= middle {
		sorted[k] = a[i]
		i++
		k++
	}

	for j <= n {
		sorted[k] = a[j]
		j++
		k++
	}

	// 정렬된 배열을 삽입
	for t := m; t <= n; t++ {
		a[t] = sorted[t]
	}
}
func mergeSort(a []int, m int, n int) {

	if m < n {
		middle := (m + n) / 2
		mergeSort(a, m, middle)
		mergeSort(a, middle+1, n)
		merge(a, m, middle, n)
	}

}

func main() {
	array := []int{7, 6, 5, 8, 3, 5, 9, 1}
	mergeSort2(array, 0, number-1)
	fmt.Println(array)
}
