package matrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	//matrix := [][]int{{1, 3}}
	fmt.Println(searchMatrix(matrix, 16))

}

func searchMatrix(matrix [][]int, target int) bool {
	line := len(matrix) - 1
	column := len(matrix[0]) - 1
	searchLine, searchColumn := 0, 0
	for searchLine <= line && searchColumn <= column {
		v := matrix[searchLine][searchColumn]
		down := getDown(matrix, searchLine, searchColumn, line)
		right := getRight(matrix, searchColumn, searchLine, column)
		if v == target {
			return true
		}
		if searchLine < line && target >= down {
			searchLine++
			continue
		}
		if searchColumn < column && target >= right {
			searchColumn++
			continue
		}
		return false
	}
	return false
}

func getDown(matrix [][]int, searchLine, searchColumn, line int) int {
	if searchLine < line {
		searchLine += 1
	}
	return matrix[searchLine][searchColumn]
}

func getRight(matrix [][]int, searchColumn, searchLine, column int) int {
	if searchColumn < column {
		searchColumn += 1
	}
	return matrix[searchLine][searchColumn]
}
