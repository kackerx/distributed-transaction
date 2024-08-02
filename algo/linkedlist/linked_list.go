package linkedlist

import "fmt"

type Linkedlist struct {
	head *node
	size int
}

func NewLinkedlist(s ...int) *Linkedlist {
	l := &Linkedlist{new(node), len(s)}

	for _, item := range s {
		l.addNode(item)
	}

	return l
}

func (l *Linkedlist) addNode(e int) {
	if l.head == nil {
		return
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &node{e, nil}
}

func (l *Linkedlist) Reverse() {
	dummy := new(node)

	cur := l.head.next
	for cur != nil {
		next := cur.next
		cur.next = dummy.next
		dummy.next = cur
		cur = next
	}

	l.head = dummy
}

func (l *Linkedlist) String() string {
	var res string

	cur := l.head.next
	for cur != nil {
		res = fmt.Sprintf(`%s-%d`, res, cur.data)
		cur = cur.next
	}

	return res
}

// ------- node

type node struct {
	data int
	next *node
}

func newNode(data int, next *node) *node {
	return &node{data, next}
}

func (n *node) setNext(e int) {
	n.next = &node{e, nil}
}
