package list

import (
	"fmt"
	"testing"
)

type MyLinkedList struct {
	Val  int
	Head *Node
	Tail *Node
}

func MyLinkedListConstructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	search := 0
	node := this.Head
	for node != nil {
		if index == search {
			return node.Val
		}
		search++
		node = node.Next
	}
	return 0
}

func (this *MyLinkedList) getIndexPreNode(index int) *Node {
	search := 0
	node := this.Head
	for node != nil {
		if search+1 == index {
			break
		}
		search++
		node = node.Next
	}
	return node
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{}
	node.Val = val
	if this.Head == nil {
		this.Head = node
		this.Tail = node
	} else {
		node.Next = this.Head
		this.Head = node
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{}
	node.Val = val
	if this.Tail == nil {
		this.Head = node
		this.Tail = node
	} else {
		node := &Node{}
		node.Val = val
		this.Tail.Next = node
		this.Tail = node
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//添加的是不是头
	if index == 0 {
		node := &Node{Val: val}
		node.Next = this.Head
		this.Head = node
	} else {
		node := this.getIndexPreNode(index)
		newNode := &Node{Val: val}
		newNode.Next = node.Next
		node.Next = newNode
		if node == this.Tail {
			this.Tail = newNode
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//删除的是不是头
	if index == 0 {
		node := this.Head
		this.Head = this.Head.Next
		node.Next = nil
	} else {
		node := this.getIndexPreNode(index)
		if node.Next == nil {
			return
		}
		if node.Next == this.Tail {
			node.Next = nil
			this.Tail = node
		} else {
			cur := node.Next
			node.Next = node.Next.Next
			cur.Next = nil
		}
	}
}

func (this *MyLinkedList) PrintFromHead() {
	node := this.Head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestMyLinkedList(t *testing.T) {
	myLinkedList := MyLinkedListConstructor()
	myLinkedList.AddAtHead(2)
	myLinkedList.DeleteAtIndex(1)
	myLinkedList.AddAtHead(2)
	myLinkedList.AddAtHead(7)
	myLinkedList.AddAtHead(3)
	myLinkedList.AddAtHead(2)
	//myLinkedList.AddAtIndex(3, 0)
	//myLinkedList.PrintFromHead()
	myLinkedList.AddAtTail(5)
	myLinkedList.DeleteAtIndex(6)
	myLinkedList.DeleteAtIndex(4)
	//myLinkedList.AddAtHead(6)
	//myLinkedList.AddAtTail(4)
	//fmt.Println(myLinkedList.Get(4))
}
