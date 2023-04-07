package search

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 1))
}

func searchRange(nums []int, target int) []int {
	first := -1
	last := -1
	if len(nums) == 0 {
		return []int{first, last}
	}
	findFirst := false
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			continue
		}
		if !findFirst {
			findFirst = true
			first = i
		}
		last = i
	}
	return []int{first, last}
}

func TestSumBase(t *testing.T) {
	sumBase(34, 6)
}

func sumBase(n int, k int) int {
	res := strconv.FormatInt(int64(n), k)
	resSlice := strings.Split(res, "")
	result := 0
	for i := 0; i < len(resSlice); i++ {
		b := resSlice[i]
		i, _ := strconv.Atoi(b)
		result += i
	}
	return result
}
