package main

func bubbleSort(raw []int) []int {
	copied := make([]int, len(raw))
	copy(copied, raw)

	for i := 0; i < len(copied)-1; i++ {
		for j := 0; j < len(copied)-1-i; j++ {
			if copied[j] > copied[j+1] {
				swap(&copied[j], &copied[j+1])
			}
		}
	}

	return copied
}
