package str

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/valid-palindrome/
func TestValidPalindromeString(t *testing.T) {
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindrome("aba"))
	fmt.Println(isPalindrome("Aba"))
	fmt.Println(isPalindrome("A b ... a"))
}

func isPalindrome(s string) bool {

	buf := bytes.Buffer{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) || (b >= 48 && b <= 57) {
			buf.WriteString(strings.ToLower(string(b)))
		}
	}
	s = buf.String()
	left := 0
	right := len(s) - 1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
