package main

import "fmt"

func zeroStripping(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	// Check if the first row initially contains a zero.
	firstRowHasZero := false
	for c := 0; c < n; c++ {
		if matrix[0][c] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Check if the first column initially contains a zero
	firstColHasZero := false
	for r := 0; r < m; r++ {
		if matrix[r][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// Use the first row and column as markers.
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}

	// Apply the zeros to the rest of the matrix according to the markers

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// If the first row originally had a zero, set it all to zero
	if firstRowHasZero {
		for c := 0; c < n; c++ {
			matrix[0][c] = 0
		}
	}

	// If the first column originally had a zero, set it all to zero.

	if firstColHasZero {
		for r := 0; r < m; r++ {
			matrix[r][0] = 0
		}
	}

}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func exec_zero_stripping() {
	matrix1 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 0, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 0},
	}

	zeroStripping(matrix1)
	printMatrix(matrix1)
}
