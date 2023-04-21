package list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func createList(size int) *Node {
	if size == 0 {
		return nil
	}
	node := &Node{}
	head := node
	for i := 0; i < size; i++ {
		node.Val = i + 1
		if i == size-1 {
			node.Next = nil
		} else {
			node.Next = &Node{}
		}
		node = node.Next
	}
	return head
}

func createListBySlice(data []int) *Node {
	if len(data) == 0 {
		return nil
	}
	node := &Node{}
	head := node
	for i := 0; i < len(data); i++ {
		node.Val = data[i]
		if i == len(data)-1 {
			node.Next = nil
		} else {
			node.Next = &Node{}
		}
		node = node.Next
	}
	return head
}

func PrintNodeList(node *Node) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
