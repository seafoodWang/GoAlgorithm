package arr

import (
	"fmt"
	"testing"
)

func TestKLengthApart(t *testing.T) {
	k := 3
	//nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	nums := []int{0, 0, 0, 0, 0, 0, 0}
	fmt.Println(kLengthApart(nums, k))
}

func kLengthApart(nums []int, k int) bool {
	lastOneIndex := -1
	for i, v := range nums {
		if v == 0 {
			continue
		}
		if lastOneIndex == -1 {
			lastOneIndex = i
			continue
		}
		if i-lastOneIndex+1 < k {
			return false
		}
	}
	if lastOneIndex == -1 {
		return false
	}
	return true
}
