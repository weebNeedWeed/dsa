package main

import "testing"

func TestMaxSubArrayOfSizeK(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{2, 3, 4, 1, 5}, 2, 7},
		{[]int{1, 1, 1, 1, 1}, 2, 2},
		{[]int{10, 20, 30, 40, 50}, 3, 120},
		{[]int{-1, -2, -3, -4, -5}, 2, -3},
		{[]int{5, 5, 5, 5, 5}, 1, 5},
		{[]int{1, 2, 3, 4, 5, 6}, 6, 21},
		{[]int{3, -1, 4, -1, 5, -9}, 3, 8},
	}

	for _, test := range tests {
		result := maxSubArrayOfSizeK(test.arr, test.k)
		if result != test.expected {
			t.Errorf("For input arr=%v and k=%d, expected %d but got %d", test.arr, test.k, test.expected, result)
		}
	}
}
