package main

func selectionSort(raw []int) []int {
	copied := make([]int, len(raw))
	copy(copied, raw)

	for i := 0; i < len(copied)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(copied); j++ {
			if copied[minIndex] > copied[j] {
				minIndex = j
			}
		}

		temp := copied[i]
		copied[i] = copied[minIndex]
		copied[minIndex] = temp
	}

	return copied
}
