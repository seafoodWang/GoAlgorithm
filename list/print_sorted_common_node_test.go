package list

import (
	"fmt"
	"testing"
)

func TestPrintSortedCommonNode(t *testing.T) {

	head1 := createListBySlice([]int{1, 2, 5, 7, 9, 10})
	head2 := createListBySlice([]int{9, 10, 11, 15})
	PrintCommonNode(head1, head2)
}

func PrintCommonNode(head1, head2 *Node) {
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			head1 = head1.Next
		} else if head2.Val < head1.Val {
			head2 = head2.Next
		} else {
			fmt.Println(head1.Val)
			head1 = head1.Next
			head2 = head2.Next
		}
	}
}
