package list

import (
	"fmt"
	"testing"
)

func TestSortColor(t *testing.T) {
	sortColors([]int{2, 0, 2, 1, 1, 0})
	//PrintNodeList(head)
}

func sortColors(nums []int) {
	red, white := 0, 1
	var redHead *Node
	var redTail *Node
	var whiteHead *Node
	var whiteTail *Node
	var blueHead *Node
	var blueTail *Node

	for i := 0; i < len(nums); i++ {
		data := nums[i]
		node := &Node{Val: data}
		if data == red {
			if redHead == nil {
				redHead = node
				redTail = node
			} else {
				redTail.Next = node
				redTail = redTail.Next
			}
		} else if data == white {
			if whiteHead == nil {
				whiteHead = node
				whiteTail = node
			} else {
				whiteTail.Next = node
				whiteTail = whiteTail.Next
			}
		} else {
			if blueHead == nil {
				blueHead = node
				blueTail = node
			} else {
				blueTail.Next = node
				blueTail = blueTail.Next
			}
		}
	}
	head := getNewHead(redHead, redTail, whiteHead)
	head = getNewHead(head, whiteTail, blueHead)
	nums = []int{}
	node := head
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	fmt.Println(nums)
}

func getNewHead(h1, t1, h2 *Node) *Node {
	if h1 == nil {
		return h2
	} else {
		if h2 != nil {
			t1.Next = h2
		}
	}

	return h1
}
