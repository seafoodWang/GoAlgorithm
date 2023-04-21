package list

import "testing"

// leetcode 18 https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func TestDeleteListNode(t *testing.T) {

	head := createList(10)
	head = deleteNodeByValue(head, 1)
	PrintNodeList(head)
}

func deleteNodeByValue(node *Node, target int) *Node {
	if node == nil {
		return nil
	}
	//第一个元素
	if node.Val == target {
		node = node.Next
		return node
	}
	slow := node
	quick := node
	quick = quick.Next
	for quick != nil {
		if quick.Val == target {
			slow.Next = slow.Next.Next
			break
		}
		slow = slow.Next
		quick = quick.Next
	}
	return node
}
