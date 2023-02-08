package linkedlist

import "errors"

type Node struct {
	next     *Node
	previous *Node
	Value    interface{}
}

type List struct {
	first *Node
	last  *Node
}

func NewList(args ...interface{}) *List {
	var list *List
	for _, v := range args {
		list.Push(v)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v, previous: nil}
	if l.first == nil {
		l.first = newNode
		l.last = newNode
		newNode.next = nil
	} else {
		newNode.next = l.first
		l.first = newNode
		newNode.next.previous = newNode
	}
}

func (l *List) Push(v interface{}) {
	if l.last == nil {
		l.Unshift(v)
		return
	}
	l.last.next = &Node{Value: v, previous: l.last, next: nil}
	l.last = l.last.next
}

func (l *List) Shift() (interface{}, error) {
	if l.first == nil {
		return nil, errors.New("")
	}
	shiftedVal := l.first.Value
	l.first = l.first.next
	l.first.previous = nil
	return shiftedVal, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.last == nil {
		return nil, errors.New("")
	}
	poppedVal := l.last.Value
	l.last = l.last.previous
	l.last.next = nil
	return poppedVal, nil
}

func recur(node *Node) *Node {
	node.next = node.previous
	node.next.previous = node.next
	if node.previous == nil {
		return node
	}
	return recur(node.previous)
}

func (l *List) Reverse() {
	recur(l.last)
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
