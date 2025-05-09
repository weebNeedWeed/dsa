package main

import "fmt"

func main() {
	demo := pascalTriangle(5)
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%-2d", demo[i][j])
		}
		fmt.Println()
	}
}

func pascalTriangle(n int) [][]int {
	pt := make([][]int, n)

	for i := 0; i < n; i++ {
		pt[i] = make([]int, i+1)
	}

	for i := 0; i < n; i++ {
		pt[i][0] = 1 // let the first column to be 1

		for j := 1; j < i; j++ {
			pt[i][j] = pt[i-1][j-1] + pt[i-1][j]
		}

		pt[i][i] = 1 // let the last column of each row to be 1
	}

	return pt
}
