package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		x    int
		want bool
	}{
		{name: "empty array", arr: []int{}, x: 10, want: false},
		{name: "single element exists", arr: []int{5}, x: 5, want: true},
		{name: "single element not exists", arr: []int{7}, x: 5, want: false},
		{name: "found in middle", arr: []int{1, 3, 5, 7, 9}, x: 5, want: true},
		{name: "not found", arr: []int{1, 3, 5, 7, 9}, x: 6, want: false},
		{name: "first element", arr: []int{1, 3, 5, 7, 9}, x: 1, want: true},
		{name: "last element", arr: []int{1, 3, 5, 7, 9}, x: 9, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.arr, tt.x); got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %v; want %v", tt.arr, tt.x, got, tt.want)
			}
		})
	}
}
