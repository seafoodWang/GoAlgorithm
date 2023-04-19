package arr

import (
	"bytes"
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/integer-to-roman/description/
func TestIntegerToRoman(t *testing.T) {
	fmt.Println(intToRoman(1994))

}

func intToRoman(num int) string {
	// num >= 1 && num <= 3999
	buf := bytes.Buffer{}
	k := num / 1000
	kStr := getStr(k, "M", "", "")
	buf.WriteString(kStr)
	h := num / 100 % 10
	hStr := getStr(h, "C", "D", "M")
	buf.WriteString(hStr)
	t := num / 10 % 10
	tStr := getStr(t, "X", "L", "C")
	buf.WriteString(tStr)
	o := num % 10
	oStr := getStr(o, "I", "V", "X")
	buf.WriteString(oStr)
	return buf.String()
}

func getStr(n int, left, mid, right string) string {
	buf := bytes.Buffer{}
	if n < 4 {
		for i := 0; i < n; i++ {
			buf.WriteString(left)
		}
	} else if n == 4 {
		buf.WriteString(left)
		buf.WriteString(mid)
	} else if n == 5 {
		buf.WriteString(mid)
	} else if n > 5 && n < 9 {
		buf.WriteString(mid)
		for i := 0; i < n-5; i++ {
			buf.WriteString(left)
		}
	} else {
		buf.WriteString(left)
		buf.WriteString(right)
	}
	return buf.String()
}
