package arr

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddDigits(t *testing.T) {
	fmt.Println(addDigits(0))
}

func addDigits(num int) int {
	str := strconv.Itoa(num)
	for len(str) != 1 {
		n := 0
		for i := 0; i < len(str); i++ {
			a, _ := strconv.Atoi(string(str[i]))
			n += a
		}
		str = strconv.Itoa(n)
	}
	d, _ := strconv.Atoi(str)
	return d
}
