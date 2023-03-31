package list

import (
	"testing"
)

func TestMergeTwoSortedList(t *testing.T) {

	head1 := createListBySlice([]int{})
	head2 := createListBySlice([]int{})
	h := mergeFromHead(head1, head2)
	PrintNodeList(h)
}

func mergeFromHead(head1, head2 *Node) *Node {
	//head1为空，返回head2
	if head1 == nil {
		return head2
	}
	//head2为空，返回head1
	if head2 == nil {
		return head1
	}
	//定义两个指针新头，尾
	var head *Node
	var tail *Node
	//两个链表存在长度不一样的情况，当有一个已经遍历完就不继续
	for head1 != nil && head2 != nil {
		//定义临时node存信息
		var node *Node
		//取当前两个节点的最小值
		if head1.Val < head2.Val {
			node = head1
			head1 = head1.Next
		} else {
			node = head2
			head2 = head2.Next
		}
		//第一次给head和tail赋值
		if head == nil {
			head = node
			tail = node
		} else {
			//非第一次，移动tail指针
			tail.Next = node
			tail = tail.Next
		}
	}
	//拼接未遍历完的数据到tail后
	if head1 != nil {
		//拼接到node后
		tail.Next = head1
	}
	if head2 != nil {
		//拼接到node后
		tail.Next = head2
	}
	return head
}
