package main

func quickSort(raw []int) []int {
	copied := make([]int, len(raw))
	copy(copied, raw)

	quickSortUtil(copied, 0, len(copied)-1)

	return copied
}

func quickSortUtil(raw []int, low, high int) {
	if low >= high {
		return
	}

	pivot := partition(raw, low, high)

	quickSortUtil(raw, low, pivot-1)
	quickSortUtil(raw, pivot+1, high)
}

func partition(raw []int, low, high int) int {
	pivot := raw[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if raw[j] < pivot {
			i++
			swap(&raw[i], &raw[j])
		}
	}

	i++
	swap(&raw[i], &raw[high])

	return i
}
