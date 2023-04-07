package arr

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/kLl5u1/description/
func TestSortedTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {
	//k->numbers value, v->numbers index
	dict := make(map[int]int)
	result := []int{0, 1}
	for k, v := range numbers {
		diff := target - v
		index, ok := dict[diff]
		if ok {
			result[0] = index
			result[1] = k
			break
		}
		dict[v] = k
	}
	return result
}
