package list

import (
	"fmt"
	"testing"
)

func TestListHasCycle(t *testing.T) {
	head := createList(10)

	fmt.Println(hasCycle(head))

}

func hasCycle(head *Node) bool {
	//头节点为空或者头节点下个节点为空
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		//慢指针每次走1步
		slow = slow.Next
		//快指针每次走两步
		fast = fast.Next.Next
		//比较快慢指针是否相等
		if slow == fast {
			return true
		}
	}
	return false
}
