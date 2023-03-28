package linkedlist

import "errors"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func New(elements []int) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	n := &Node{element, nil}
	if l.tail == nil {
		l.head = n
		l.tail = l.head
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

var errEmptyList = errors.New("empty list")

func (l *List) Pop() (v int, e error) {
	if l.tail == nil {
		return 0, errEmptyList
	}
	l.size--
	v = l.tail.value
	if l.head == l.tail {
		l.head, l.tail = nil, nil
		return
	}
	n := l.head
	for n.next != l.tail {
		n = n.next
	}
	n.next, l.tail = nil, n

	return
}

func (l *List) Array() []int {
	arr := make([]int, l.Size())
	if l.head == nil {
		return arr
	}
	n := l.head
	for i := 0; n != nil; i++ {
		arr[i] = n.value
		n = n.next
	}
	return arr
}

func (l *List) Reverse() *List {
	r := l.Array()
	for i, j := 0, l.size-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return New(r)
}
