package list

import (
	"fmt"
	"testing"
)

func TestSearchInsertIndex(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	size := len(nums)
	left := 0
	right := size - 1
	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
