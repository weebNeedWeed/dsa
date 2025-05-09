package main

func insertionSort(raw []int) []int {
	copied := make([]int, len(raw))
	copy(copied, raw)

	for i := 1; i < len(copied); i++ {
		x := copied[i]
		pos := i - 1

		for pos >= 0 && copied[pos] > x {
			copied[pos+1] = copied[pos]
			pos--
		}

		copied[pos+1] = x
	}

	return copied
}
