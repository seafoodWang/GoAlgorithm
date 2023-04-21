package str

import (
	"fmt"
	"testing"
)

func TestPeekInMountains(t *testing.T) {
	fmt.Println(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
}

func peakIndexInMountainArray(arr []int) int {
	//找这个数组的最大值
	left := 0
	right := len(arr) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		midVal := arr[mid]
		midLeft := arr[mid-1]
		midRight := arr[mid+1]
		if midLeft < midVal && midVal > midRight {
			break
		}
		if midLeft > midVal {
			right = mid
		} else {
			left = mid
		}
	}
	return mid
}
