package list

import (
	"fmt"
	"testing"
)

//leetcode-19 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func TestRemoveNthNodeFromEndOfList1st(t *testing.T) {

	head := createList(10)
	//aimIndex := 5
	node := head
	iteratorTimes := 0
	for node.Next != nil {
		iteratorTimes++
	}
	fmt.Printf("%d", iteratorTimes)
}

func TestRemoveNthNodeFromEndOfList2nd(t *testing.T) {

}
