package str

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBullAndCows(t *testing.T) {
	fmt.Println(getHint("1122", "0001"))
}

func getHint(secret string, guess string) string {
	str1, bullIndex := getBulls(secret, guess)
	str2 := getCows(secret, guess, bullIndex)
	return str1 + str2
}

func getBulls(secret string, guess string) (string, map[int]struct{}) {
	bullIndex := make(map[int]struct{}, 0)
	cnt := 0
	for i := 0; i < len(secret); i++ {
		s := secret[i]
		g := guess[i]
		if s == g {
			cnt++
			bullIndex[i] = struct{}{}
		}
	}
	return strconv.Itoa(cnt) + "A", bullIndex
}

func getCows(secret string, guess string, bullIndex map[int]struct{}) string {
	sMap := make(map[int]int)
	gMap := make(map[int]int)
	for i := 0; i < len(secret); i++ {
		_, ok := bullIndex[i]
		if ok {
			continue
		}
		s := int(secret[i])
		sValue := sMap[s]
		sValue++
		sMap[s] = sValue

		g := int(guess[i])
		gValue := gMap[g]
		gValue++
		gMap[g] = gValue
	}
	total := 0
	for k, sCount := range sMap {
		gCount, ok := gMap[k]
		if !ok {
			continue
		}
		min := sCount
		if gCount < sCount {
			min = gCount
		}
		total += min
	}
	return strconv.Itoa(total) + "B"
}
