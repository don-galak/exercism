package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func New(elements []int) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (l *List) Size() int {
	c := 0
	node := l.head
	for node != nil {
		c++
		node = node.next
	}
	return c
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
}

var errEmptyList = errors.New("empty list")

func (l *List) Pop() (v int, e error) {
	if l.tail == nil {
		return 0, errEmptyList
	}

	v = l.tail.value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}
	n := l.head
	for n.next != l.tail {
		n = n.next
		println(n.value)
	}
	l.tail = n

	fmt.Printf("list: %v\n", l.Array())
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
	panic("Please implement the Reverse function")
}
