package main

func selectionSort(data []int) {
	for i, _ := range data {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[min] > data[j] {
				min = j
			}
		}
		data[min], data[i] = data[i], data[min]
	}
}
