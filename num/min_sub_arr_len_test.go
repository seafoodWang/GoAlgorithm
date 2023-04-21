package num

import (
	"fmt"
	"testing"
)

func TestMinArrSubLen(t *testing.T) {
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}

func minSubArrayLen(target int, nums []int) int {
	//window wide
	windowWide := 1
	for windowWide <= len(nums) {
		//已windowWide宽度进行滑动
		for i := 0; i < len(nums)-windowWide+1; i++ {
			windowNum := 0
			//滑动窗口总值
			for j := i; j < i+windowWide; j++ {
				windowNum += nums[j]
			}
			if windowNum >= target {
				return windowWide
			}
		}
		windowWide++
	}
	return 0
}
