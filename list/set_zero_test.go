package list

import (
	"fmt"
	"testing"
)

//https://leetcode.cn/problems/set-matrix-zeroes/description/

func TestSetMatrixZeros(t *testing.T) {
	nums := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeros(nums)
	fmt.Println(nums)
}

func setZeros(matrix [][]int) {
	//标记数据
	lineZero := make(map[int]struct{})
	columnZero := make(map[int]struct{})
	for line := 0; line < len(matrix); line++ {
		lineData := matrix[line]
		for column := 0; column < len(lineData); column++ {
			if lineData[column] == 0 {
				lineZero[line] = struct{}{}
				columnZero[column] = struct{}{}
			}
		}
	}
	//清除
	for line := 0; line < len(matrix); line++ {
		lineData := matrix[line]
		for column := 0; column < len(lineData); column++ {
			lineSweep := needSweep(line, lineZero)
			columnSweep := needSweep(column, columnZero)
			if lineSweep || columnSweep {
				lineData[column] = 0
			}
		}
	}
}

func needSweep(index int, data map[int]struct{}) bool {
	if _, ok := data[index]; ok {
		return true
	}
	return false
}
