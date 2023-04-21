package arr

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons
func TestFindMinArrowShot(t *testing.T) {
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(findMinArrowShots(points))
}

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	end := points[0][1]
	for _, point := range points {
		if point[0] > end {
			arrows++
			end = point[1]
		}
	}

	return arrows
}
