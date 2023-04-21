package num

import (
	"fmt"
	"testing"
)

func TestBitWise(t *testing.T) {
	fmt.Println(rangeBitwiseAnd(100, 2147483644))
}

func rangeBitwiseAnd(left int, right int) int {
	n := left
	for left <= right {
		n &= left
		left++
	}
	return n
}
