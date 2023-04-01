package list

import (
	"testing"
)

func TestLinkedListReverse(t *testing.T) {

	head := createList(10)
	h := reverse(head)
	PrintNodeList(h)
}

func reverse(head *Node) *Node {
	if head == nil {
		return nil
	}
	var node *Node
	for head != nil {
		//原链表的头节点给cur
		cur := head
		//原链表头指针后移
		head = head.Next
		//cur的Next给node
		cur.Next = node
		//新链表头向前移动
		node = cur
	}
	return node
}
