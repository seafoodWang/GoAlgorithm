package arr

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/merge-intervals/
func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{{2, 4}, {1, 5}, {8, 10}, {15, 19}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	mergeSlice := make([][]int, 0)
	for _, v := range intervals {
		if len(mergeSlice) == 0 || v[0] > mergeSlice[len(mergeSlice)-1][1] {
			mergeSlice = append(mergeSlice, v)
		} else {
			maxValue := max(v[1], mergeSlice[len(mergeSlice)-1][1])
			mergeSlice[len(mergeSlice)-1][1] = maxValue
		}
	}
	return mergeSlice
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
