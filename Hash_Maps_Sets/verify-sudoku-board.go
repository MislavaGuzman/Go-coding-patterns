package main

import "fmt"

func verifySudokuBoard(board [9][9]int) bool {
	var rowSets [9]map[int]bool
	var columnSets [9]map[int]bool
	var subgridSets [3][3]map[int]bool

	for i := 0; i < 9; i++ {
		rowSets[i] = make(map[int]bool)
		columnSets[i] = make(map[int]bool)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subgridSets[i][j] = make(map[int]bool)
		}
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			num := board[r][c]
			if num == 0 {
				continue
			}

			subgridRow, subgridCol := r/3, c/3

			if rowSets[r][num] {
				return false
			}

			if columnSets[c][num] {
				return false
			}

			if subgridSets[subgridRow][subgridCol][num] {
				return false
			}

			rowSets[r][num] = true
			columnSets[c][num] = true
			subgridSets[subgridRow][subgridCol][num] = true
		}
	}

	return true
}

func run_verifySudokuBoard() {
	validBoard := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	invalidBoard := [9][9]int{
		{5, 5, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println(verifySudokuBoard(validBoard))
	fmt.Println(verifySudokuBoard(invalidBoard))
}
