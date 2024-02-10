package main

func InsertionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			} else {
				break
			}
		}
	}
}
