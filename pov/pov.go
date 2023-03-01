package pov

type Tree struct {
	value    string
	parent   *Tree
	children []*Tree
	// Add the needed fields here
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	root := &Tree{value: value, children: children}
	for _, ch := range children {
		ch.parent = root
	}
	return root
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions
func findNode(node *Tree, from string) *Tree {
	if node.value == from {
		return node
	}
	for _, ch := range node.children {
		l := findNode(ch, from)
		if l != nil && l.value == from {
			return l
		}
	}
	return nil
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	if tr.value == from {
		return tr
	}
	node := findNode(tr, from)
	if node == nil {
		return nil
	}

	parent := node.parent
	var index int
	for i, ch := range parent.children {
		if ch.value == from {
			index = i
			break
		}
	}

	parent.children[index] = parent.children[len(parent.children)-1]
	parent.children = parent.children[:len(parent.children)-1]
	if parent.parent != nil {
		println("PARENT PARENT: ", parent.parent.String())
		// parent.children = []*Tree{parent.parent}
	}
	parent.parent = node
	node.parent = nil
	node.children = append(node.children, parent)

	return node
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	panic("Please implement this function")
}
