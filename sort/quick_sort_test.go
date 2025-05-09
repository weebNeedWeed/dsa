package main

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"zero elements", []int{}, []int{}},
		{"normal array with negative and non-negative numbers", []int{2, -1, -99, 0, 4}, []int{-99, -1, 0, 2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSort(tt.input)

			if len(got) != len(tt.want) {
				t.Errorf("error when sorting array; got %v, want %v", got, tt.want)
			}

			for i, _ := range got {
				if got[i] != tt.want[i] {
					t.Errorf("error when sorting array; got %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
