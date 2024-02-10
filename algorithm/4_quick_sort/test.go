package main

func QuickSortDesc(data []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := start
	i := start + 1
	j := end

	for i <= j {
		for i <= end && data[i] >= data[pivot] {
			i++
		}
		for j >= start && data[j] <= data[pivot] {
			j--
		}
		if i > j {
			data[j], data[pivot] = data[pivot], data[j]
		} else {
			data[i], data[j] = data[j], data[i]
		}
	}

	QuickSortDesc(data, start, j-1)
	QuickSortDesc(data, j+1, end)
}
