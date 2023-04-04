package list

import "testing"

func TestDeleteDuplicateNode(t *testing.T) {
	head := createListBySlice([]int{1, 1, 2, 3, 4, 4, 5, 5})
	head = deleteDuplicates(head)
	PrintNodeList(head)
}

func deleteDuplicates(head *Node) *Node {
	if head == nil {
		return nil
	}
	node := head
	var prev *Node
	for node != nil {
		if prev == nil || prev.Val != node.Val {
			prev = node
		} else {
			prev.Next = node.Next
		}
		node = node.Next
	}
	return head
}
