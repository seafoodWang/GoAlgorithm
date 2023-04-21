package list

import (
	"testing"
)

//leetcode-19 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func TestRemoveNthNodeFromEndOfList1st(t *testing.T) {
	size := 10
	head := createList(size)
	lastIndex := 10
	head = solution1(size, lastIndex, head)
	PrintNodeList(head)
}

func solution1(size, lastIndex int, head *Node) *Node {
	if size == lastIndex {
		head = head.Next
		return head
	}
	if lastIndex == 0 {
		return head
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
	return head
}

// two ptr
func TestRemoveNthNodeFromEndOfList2nd(t *testing.T) {
	size := 10
	head := createList(size)
	lastIndex := 10
	head = solution2(lastIndex, head)
	PrintNodeList(head)
}

func solution2(lastIndex int, head *Node) *Node {
	return nil
}
