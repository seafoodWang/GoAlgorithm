package list

import (
	"fmt"
	"testing"
)

func TestValidSudoku(t *testing.T) {
	board := getData1()
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	//行
	for line := 0; line < 9; line++ {
		//列
		for column := 0; column < 9; column++ {
			nineHouseData := getNineHouse(line, column, board)
			lineData := getLineData(line, board)
			columnData := getColumnData(column, board)
			columnRes := checkData(columnData)
			lineRes := checkData(lineData)
			nineHouseRes := checkData(nineHouseData)
			isValid := columnRes && lineRes && nineHouseRes
			if !isValid {
				return false
			}
		}
	}
	return true
}

func getLineData(line int, board [][]byte) []byte {
	res := make([]byte, 9, 9)
	for column := 0; column < 9; column++ {
		b := board[column][line]
		res[column] = b
	}
	return res
}

func checkData(data []byte) bool {
	mapping := make(map[string]struct{})
	for _, val := range data {
		s := string(val)
		if s == "." {
			continue
		}
		_, ok := mapping[s]
		if ok {
			return false
		} else {
			mapping[s] = struct{}{}
		}
	}
	return true
}

func getColumnData(column int, board [][]byte) []byte {
	res := make([]byte, 9, 9)
	for line := 0; line < 9; line++ {
		b := board[column][line]
		res[line] = b
	}
	return res
}

func getNineHouse(line, column int, board [][]byte) []byte {
	baseLine := (line / 3) * 3
	baseColumn := (column / 3) * 3
	res := make([]byte, 9)
	t := 0
	for i := 0; i < 3; i++ {
		line := baseLine + i
		for j := 0; j < 3; j++ {
			column := baseColumn + j
			res[t] = board[column][line]
			t++
		}
	}
	return res
}

func getData1() [][]byte {
	b1 := []byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'}
	b2 := []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	b3 := []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	b4 := []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	b5 := []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	b6 := []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	b7 := []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	b8 := []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	b9 := []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	return [][]byte{b1, b2, b3, b4, b5, b6, b7, b8, b9}
}
