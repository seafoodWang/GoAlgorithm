package arr

import (
	"fmt"
	"testing"
)

func TestBulbSwitches(t *testing.T) {
	fmt.Println(bulbSwitch(5))
}

func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = 1
	}
	start := 2
	for i := start; i <= n; i++ {
		for j := 0; j < n; j++ {
			if (j+1)%i == 0 {
				r[j] = 0
			}
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if r[i] == 1 {
			cnt++
		}
	}
	return cnt
}
