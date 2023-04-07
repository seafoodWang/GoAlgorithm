package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	//s := []int{1,2,3,4,4}
	s := []int{1, 2, 3, 2, 2}
	head := createListBySlice(s)
	res := isPalindrome1(head)
	fmt.Println(res)
}

func isPalindrome1(head *Node) bool {

	stack := make([]int, 0)
	node := head
	for node != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}
	node = head
	for node != nil {
		l := len(stack)
		v := stack[l-1]
		if node.Val != v {
			return false
		}
		stack = stack[:l-1]
		node = node.Next
	}
	return true
}
