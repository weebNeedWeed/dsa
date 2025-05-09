package main

import "fmt"

func main() {
	fmt.Println("Enter n:")
	var n int
	fmt.Scanf("%d", &n)

	arr := pascalTriangle(n)

	for _, row := range arr {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}

func pascalTriangle(row int) [][]int {
	res := make([][]int, row, row)

	for i, _ := range res {
		res[i] = make([]int, i+1)
	}

	for i := 0; i < row; i++ {
		res[i][0] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}

		res[i][i] = 1
	}

	return res
}
