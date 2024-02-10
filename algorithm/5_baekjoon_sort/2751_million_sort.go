package main

func quickSort(data []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := start
	i := start + 1
	j := end

	for i <= j {
		for i <= end && data[i] <= data[pivot] {
			i++
		}

		for j > start && data[j] >= data[pivot] {
			j--
		}

		if i > j {
			data[pivot], data[j] = data[j], data[pivot]
		} else {
			data[i], data[j] = data[j], data[i]
		}
	}

	quickSort(data, start, j-1)
	quickSort(data, j+1, end)

}
