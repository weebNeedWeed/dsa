package main

func mergeSort(raw []int) []int {
	copied := make([]int, len(raw))
	copy(copied, raw)

	mergeSortUtil(copied)

	return copied
}

func mergeSortUtil(raw []int) {
	n := len(raw)
	if n <= 1 {
		return
	}

	mid := n / 2
	leftArr := make([]int, mid)
	rightArr := make([]int, n-mid)

	copy(leftArr, raw[:mid])
	copy(rightArr, raw[mid:])

	mergeSortUtil(leftArr)
	mergeSortUtil(rightArr)

	merge(leftArr, rightArr, raw)
}

func merge(left, right, output []int) {
	p1, p2 := 0, 0
	write := 0
	for p1 < len(left) && p2 < len(right) {
		if left[p1] < right[p2] {
			output[write] = left[p1]
			p1++
		} else {
			output[write] = right[p2]
			p2++
		}
		write++
	}

	for p1 < len(left) {
		output[write] = left[p1]
		p1++
		write++
	}

	for p2 < len(right) {
		output[write] = right[p2]
		p2++
		write++
	}
}
