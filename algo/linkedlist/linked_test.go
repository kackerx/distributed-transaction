package linkedlist

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	l := NewLinkedlist([]int{1, 8, 7, 5, 6, 4, 9}...)

	fmt.Println(l)

	l.Reverse()

	fmt.Println(l)
}
