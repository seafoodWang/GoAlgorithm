package list

import "testing"

func TestCreateNode(t *testing.T) {
	node := createList(5)
	index := 0
	for node != nil {
		if node.Val != index+1 {
			t.Error(node.Val, "not equal index : ", index, "index value : ", index+1)
		}
		t.Log(node.Val)
		node = node.Next
		index++
	}
	node = createList(0)
	if node != nil {
		t.Error("node not nil")
	}
}
