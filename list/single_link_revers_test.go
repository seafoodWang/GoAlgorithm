package list

import "testing"

func TestLinkedListReverse(t *testing.T) {

	head := createList(10)

	node := head
	var pre *Node
	for head.Next != nil {
		cur := node
		node = node.Next
		cur.Next = pre
		pre = cur
	}
}
