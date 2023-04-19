package list

import "testing"

func TestSwapNodeInPair(t *testing.T) {
	head := createListBySlice([]int{1, 2, 3, 4, 5})
	head = swapPairs(head)
	PrintNodeList(head)
}

func swapPairs(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var h *Node
	var n *Node
	for head != nil && head.Next != nil {
		cur := head
		next := head.Next
		head = head.Next.Next
		next.Next = nil
		cur.Next = nil
		if h == nil {
			h = next
			h.Next = cur
			n = cur
		} else {
			n.Next = next
			next.Next = cur
			n = cur
		}
	}
	if head != nil {
		n.Next = head
	}
	return h
}
