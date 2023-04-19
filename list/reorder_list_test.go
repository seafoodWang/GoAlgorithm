package list

import "testing"

func TestReorderList(t *testing.T) {
	head := createListBySlice([]int{1, 2, 3, 4, 5})
	reorderList(head)
}

func reorderList(head *Node) {
	if head == nil || head.Next == nil {
		return
	}
	nodeCnt := 0
	//找中点
	var midNode *Node
	//剩下的一半翻转链表，偶数，奇数
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		nodeCnt++
	}
	mid := nodeCnt / 2
	if nodeCnt%2 == 1 {
		mid = nodeCnt/2 + 1
	}
	tmp = head
	for i := 0; i < mid; i++ {
		midNode = tmp
		tmp = tmp.Next
	}
	l2 := midNode.Next
	midNode.Next = nil
	//l1包含中点，l2不含中点
	l2 = rev(l2)
	l1 := head
	var n *Node
	var h *Node
	for l2 != nil {
		tmp1 := l1
		tmp2 := l2
		l1 = l1.Next
		l2 = l2.Next
		tmp1.Next = tmp2
		if n == nil {
			n = tmp1
			h = n
		} else {
			n.Next = tmp1
			n = tmp2
		}
	}
	PrintNodeList(h)
}

func rev(head *Node) *Node {
	var node *Node
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = node
		node = cur
	}
	return node
}
