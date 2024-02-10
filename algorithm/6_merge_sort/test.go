package main

var sorted2 [8]int

func merge2(data []int, start int, end int, middle int) {
	if start >= end {
		return
	}

	k := start
	i := start
	j := middle + 1

	for i <= middle && j <= end {
		if data[i] <= data[j] {
			sorted2[k] = data[i]
			i++
		} else {
			sorted2[k] = data[j]
			j++
		}

		k++
	}

	for i <= middle {
		sorted2[k] = data[i]
		i++
		k++
	}

	for j <= end {
		sorted2[k] = data[j]
		j++
		k++
	}

	for t := start; t <= end; t++ {
		data[t] = sorted2[t]
	}

}

func mergeSort2(data []int, start int, end int) {
	if start >= end {
		return
	}
	middle := (start + end) / 2
	mergeSort2(data, start, middle)  // 쪼개고
	mergeSort2(data, middle+1, end)  // 쪼개고
	merge2(data, start, end, middle) // 합치고
}
