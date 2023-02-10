package linkedlist

type Node struct {
	next  *Node
	prev  *Node
	Value interface{}
}

type List struct {
	first *Node
	last  *Node
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
	if l.first == nil {
		l.first = newNode
		l.last = l.first
	} else {
		newNode.next = l.first
		l.first = newNode
		newNode.next.prev = newNode
	}
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.last == nil {
		l.first = newNode
		l.last = l.first
	} else {
		l.last.next = newNode
		newNode.prev = l.last
		l.last = newNode
	}
}

func (l *List) Shift() (val interface{}, err error) {
	if l.first != nil && l.first == l.last {
		val = l.first.Value
		l.first = nil
		l.last = nil
		return
	}
	val = l.first.Value
	l.first = l.first.next
	l.first.prev = nil
	return
}

func (l *List) Pop() (val interface{}, err error) {
	if l.last != nil && l.first == l.last {
		val = l.last.Value
		l.first = nil
		l.last = nil
		return
	}
	val = l.last.Value
	l.last = l.last.prev
	l.last.next = nil
	return
}

func (l *List) Reverse() {
	if l.first != nil && l.first != l.last {
		node := l.first
		for node != nil && node != l.last {
			value, _ := l.Shift()
			l.Push(value)
			node = node.Next()
		}
		last := l.last
		l.last = l.first
		l.first = last
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
