package list

import (
	"fmt"
	"testing"
)

//leetcode-19 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func TestRemoveNthNodeFromEndOfList1st(t *testing.T) {
	size := 10
	head := createList(10)
	lastIndex := 10
	if size == lastIndex {
		head = head.Next
		return
	}
	node := head
	iteratorTimes := 0
	for node.Next != nil {
		iteratorTimes++
		node = node.Next
	}
	iteratorTimes -= lastIndex
	node = head
	for i := 0; i < iteratorTimes; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// two ptr
// TODO
func TestRemoveNthNodeFromEndOfList2nd(t *testing.T) {

}
