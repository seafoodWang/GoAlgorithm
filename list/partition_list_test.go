package list

import "testing"

// https://leetcode.cn/problems/partition-list/
func TestPartitionList(t *testing.T) {
	head := createListBySlice([]int{2, 1})
	head = partition(head, 0)
	PrintNodeList(head)
}

func partition(head *Node, x int) *Node {
	var lHead, lTail *Node
	var eHead, eTail *Node

	for head != nil {
		if head.Val < x {
			if lHead == nil {
				lHead = head
				lTail = head
			} else {
				lTail.Next = head
				lTail = lTail.Next
			}
		} else {
			if eHead == nil {
				eHead = head
				eTail = head
			} else {
				eTail.Next = head
				eTail = eTail.Next
			}
		}
		head = head.Next
	}
	if lHead == nil {
		lHead = eHead
	} else {
		lTail.Next = eHead
	}
	if eTail != nil {
		eTail.Next = nil
	}
	return lHead
}
