package linkedlist

type Node struct {
	next  *Node
	prev  *Node
	Value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, v := range args {
		list.Push(v)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		newNode.next = l.head
		l.head = newNode
		newNode.next.prev = newNode
	}
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.tail == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

func (l *List) Shift() (val interface{}, err error) {
	if l.head != nil && l.head == l.tail {
		val = l.head.Value
		l.head = nil
		l.tail = nil
		return
	}
	val = l.head.Value
	l.head = l.head.next
	l.head.prev = nil
	return
}

func (l *List) Pop() (val interface{}, err error) {
	if l.tail != nil && l.head == l.tail {
		val = l.tail.Value
		l.head = nil
		l.tail = nil
		return
	}
	val = l.tail.Value
	l.tail = l.tail.prev
	l.tail.next = nil
	return
}

func (l *List) Reverse() {
	node := l.head
	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.prev
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
