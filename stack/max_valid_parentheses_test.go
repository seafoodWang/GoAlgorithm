package stack

import (
	"fmt"
	"sort"
	"testing"
)

//https://leetcode.cn/problems/longest-valid-parentheses/

func TestMaxValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(")((())()"))
	fmt.Println(longestValidParentheses("((("))
	fmt.Println(longestValidParentheses(")))"))
	fmt.Println(longestValidParentheses("(((())))"))
}

func longestValidParentheses(s string) int {
	validCount := 0
	validSlice := []int{0}
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		d := string(s[i])
		if d == "(" {
			stack = append(stack, d)
		} else {
			isEmpty := len(stack) == 0
			if isEmpty {
				validCount = 0
				validSlice = append(validSlice, validCount)
				continue
			}
			stack = stack[:len(stack)-1]
			validCount += 2
			validSlice = append(validSlice, validCount)
		}
	}
	sort.Ints(validSlice)
	return validSlice[len(validSlice)-1]
}

func Test(t *testing.T) {
	fmt.Println(searchRange([]int{1, 2}, 2))
}

func searchRange(nums []int, target int) []int {
	first := -1
	last := -1
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{first, last}
		}
	}
	findFirst := false
	findLast := false
	for i := 0; i < len(nums)-1; i++ {
		d := nums[i]
		if d != target {
			continue
		}
		if !findFirst {
			findFirst = true
			first = i
		}
		last = i
		if nums[i+1] != target {
			findLast = true
		}
		if findFirst && findLast {
			break
		}
	}
	res := []int{first, last}
	fmt.Println(res)
	return res
}
