package main

import "fmt"

// Quick Sort - 오름차순 정렬
// 일반적으로 quick 정렬은 가장 앞의 값을 pivot으로 잡는다.
// pivot을 기준으로 작은 값은 왼쪽, 큰 값은 오른쪽으로 정렬한다.
// 즉 왼쪽에서 오른쪽으로 가면서 pivot보다 큰 값을 찾고, 오른쪽에서 왼쪽으로 가면서 pivot보다 작은 값을 찾는다.
// 그리고 두 값을 교환한다. 이 과정을 반복한다.
// 만약 둘의 위치가 엇갈린다면 작은 값과 pivot을 교환한다.

// 시간 복잡도 : O(NlogN)
// 최악의 경우 : O(N^2)
// 일반적으로는 가장 빠르다고 알려져 있다.

func quickSort(data []int, start int, end int) {
	if start >= end { // 원소가 1개인 경우 그대로 두기
		return
	}

	key := start // 키는 첫 번째 원소
	i := start + 1
	j := end

	for i <= j { // 엇갈릴 때까지 반복
		for i <= end && data[i] <= data[key] {
			i++
		}
		for j > start && data[j] >= data[key] {
			j--
		}

		if i > j { // 현재 엇갈린 상태면 키 값과 교체
			data[j], data[key] = data[key], data[j] // j는 큰걸 찾은 거니까 j와 key를 교체
		} else { // 엇갈리지 않았다면 i와 j를 교체
			data[i], data[j] = data[j], data[i]
		}

	}

	quickSort(data, start, j-1)
	quickSort(data, j+1, end)
}

var number = 10
var data = []int{9, 1, 5, 6, 4, 10, 7, 8, 3, 2}

func quickSortDesc(data []int, start int, end int) {
	if start >= end {
		return
	}

	key := start
	i := start + 1
	j := end

	for i <= j {
		for i <= end && data[i] >= data[key] { // 큰게 좌측에 와야 하니까 작은 걸 만나면 멈춘다.
			i++
		}
		for j > start && data[j] <= data[key] {
			j--
		}

		if i > j {
			data[j], data[key] = data[key], data[j]
		} else {
			data[i], data[j] = data[j], data[i]
		}
	}

	quickSortDesc(data, start, j-1)
	quickSortDesc(data, j+1, end)
}

func main() {
	QuickSortDesc(data, 0, number-1)
	for _, v := range data {
		fmt.Print(v, " ")
	}
}
