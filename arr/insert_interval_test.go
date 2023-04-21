package arr

import (
	"fmt"
	"sort"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	interval := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{4, 5}
	fmt.Println(insert(interval, newInterval))

}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergeSlice := make([][]int, 0)
	for _, interval := range intervals {
		if len(mergeSlice) == 0 || interval[0] > mergeSlice[len(mergeSlice)-1][1] {
			mergeSlice = append(mergeSlice, interval)
		} else {
			maxValue := max(interval[1], mergeSlice[len(mergeSlice)-1][1])
			mergeSlice[len(mergeSlice)-1][1] = maxValue
		}

	}
	return mergeSlice
}
