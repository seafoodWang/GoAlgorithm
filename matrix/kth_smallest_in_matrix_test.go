package matrix

import (
	"fmt"
	"testing"
)

func TestKthInMatrix(t *testing.T) {

	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(kthSmallest(matrix, 7))
	//matrix = [][]int{{-1}}
	//fmt.Println(kthSmallest(matrix, 1))
	var matrix = [][]int{{1, 2}, {1, 3}}
	fmt.Println(kthSmallest(matrix, 2))

}

func kthSmallest(matrix [][]int, k int) int {
	line := len(matrix) - 1
	column := len(matrix[0]) - 1

	searchLine := 0
	if line != 0 {
		searchLine = k / (line + 1)
	}
	searchColumn := 0
	if column != 0 {
		searchColumn = k%(column+1) - 1
		if searchColumn < 0 {
			searchColumn = 0
		}
	}

	return matrix[searchLine][searchColumn]
}
