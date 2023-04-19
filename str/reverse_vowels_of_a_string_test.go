package str

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReverseVowel(t *testing.T) {
	res := reverseVowels("hello HELLO")
	fmt.Println(res)
}

func reverseVowels(s string) string {
	//使用栈存储
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if isVowel(b) {
			stack = append(stack, b)
		}
	}
	buf := bytes.Buffer{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if isVowel(b) {
			if len(stack)-1 >= 0 {
				b = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
			}
		}
		buf.WriteByte(b)
	}
	return buf.String()
}

func isVowel(b byte) bool {
	if b == 65 || b == 97 || b == 69 || b == 101 ||
		b == 73 || b == 105 || b == 79 || b == 111 || b == 85 || b == 117 {
		return true
	}
	return false
}
