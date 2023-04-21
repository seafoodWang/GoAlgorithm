package list

import (
	"testing"
)

// https://leetcode.cn/problems/add-two-numbers-ii/
func TestTwoListAddNode(t *testing.T) {
	l1 := createListBySlice([]int{1, 0, 0, 9})
	l2 := createListBySlice([]int{9, 9, 1})
	head := addTwoNumbers(l1, l2)
	PrintNodeList(head)
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	//翻转链表l1,l2
	l1 = reverseNode(l1)
	l2 = reverseNode(l2)
	//元素相加拼接新链表，进位
	addOne := 0
	head := &Node{}
	node := head
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + addOne
		addOne = val / 10
		cur := &Node{Val: val % 10}
		node.Next = cur
		node = cur
		l1 = l1.Next
		l2 = l2.Next
	}
	var last *Node
	if l1 == nil {
		last = l2
	}
	if l2 == nil {
		last = l1
	}
	node.Next = last

	//翻转链表，去除头部0，如果只有0则不处理
	if head.Next != nil && head.Val == 0 {
		head = head.Next
	}
	head = reverseNode(head)
	return head
}

func reverseNode(head *Node) *Node {
	var node *Node
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = node
		node = cur
	}
	return node
}
