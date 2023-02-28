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
func rec(root *Tree, from string) *Tree {
	if root.value == from {
		return root
	}
	for _, ch := range root.children {
		return rec(ch, from)
	}
	return nil
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	root := rec(tr, from)
	if root == nil {
		return nil
	}
	parent := root.parent
	var index int
	for i, ch := range parent.children {
		println(ch.value)
		if ch.value == from {
			index = i
			break
		}
	}
	parent.children[index] = parent.children[len(parent.children)-1]
	parent.children = parent.children[:len(parent.children)-1]
	root.parent = nil
	root.children = append(root.children, parent)

	return root
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	panic("Please implement this function")
}
