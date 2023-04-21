package arr

import "testing"

func TestRotateArray(t *testing.T) {
	oldArray := []int{1, 2, 3, 4, 5, 6, 7}
	newArray := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(oldArray, 3)
	for i := 0; i < len(oldArray); i++ {
		if oldArray[i] != newArray[i] {
			t.Error("old not equal to new at index :", i)
			break
		}
	}
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
