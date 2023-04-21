package list

import (
	"fmt"
	"testing"
)

func TestIsPalindromeList(t *testing.T) {
	head := createListBySlice([]int{1, 2, 3, 2, 1})
	r := isPalindrome(head)
	fmt.Println(r)

	head = createListBySlice([]int{1, 2})
	r = isPalindrome(head)
	fmt.Println(r)
}

func isPalindrome(head *Node) bool {
	if head == nil {
		return true
	}
	node := head
	stack := make([]int, 0)
	for node != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}
	node = head
	index := len(stack) - 1
	for node != nil {
		if stack[index] != node.Val {
			return false
		}
		index--
		node = node.Next
	}
	return true

}
