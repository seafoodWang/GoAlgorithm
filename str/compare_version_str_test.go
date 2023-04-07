package str

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// 165 https://leetcode.cn/problems/compare-version-numbers/
func TestCompareVersion(t *testing.T) {
	fmt.Println(compareVersion("1", "1.1"))
	//fmt.Println(compareVersion("0.1", "0.1.0"))
	//fmt.Println(compareVersion("1", "0.1"))
}

func compareVersion(version1 string, version2 string) int {
	version1S := strings.Split(version1, ".")
	version2S := strings.Split(version2, ".")

	//短的补齐(.0) long-short次
	v1Length := len(version1S)
	v2Length := len(version2S)
	length := v1Length
	diff := math.Abs(float64(v1Length - v2Length))
	diffCnt := int(diff)
	if v1Length > v2Length {
		for i := 0; i < diffCnt; i++ {
			version2S = append(version2S, "0")
		}
	} else if v1Length < v2Length {
		length = v2Length
		for i := 0; i < diffCnt; i++ {
			version1S = append(version1S, "0")
		}
	}
	for i := 0; i < length; i++ {
		s1 := strings.TrimLeft(version1S[i], "0")
		s2 := strings.TrimLeft(version2S[i], "0")
		if s1 == "" {
			s1 = "0"
		}
		if s2 == "" {
			s2 = "0"
		}
		i1, _ := strconv.Atoi(s1)
		i2, _ := strconv.Atoi(s2)
		if i1 == i2 {
			continue
		}
		if i1 < i2 {
			return -1
		} else {
			return 1
		}
	}
	return 0
}
