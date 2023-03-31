package list

type Node struct {
	Val  int
	Next *Node
}

func createList(size int) *Node {
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
