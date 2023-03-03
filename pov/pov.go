package pov

type Tree struct {
	value    string
	children []*Tree
}

func New(value string, children ...*Tree) *Tree { return &Tree{value: value, children: children} }

func (tr *Tree) Value() string     { return tr.value }
func (tr *Tree) Children() []*Tree { return tr.children }
func (tr *Tree) createPath(from string, path []*Tree) []*Tree {
	path = append(path, tr)
	if tr.value == from {
		return path
	}
	for _, ch := range tr.children {
		if p := ch.createPath(from, path); p != nil {
			return p
		}
	}
	return nil
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
// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	path := tr.createPath(from, make([]*Tree, 0))
	pathLength := len(path)
	if pathLength == 0 {
		return nil
	}

	for i := pathLength - 1; i > 0; i-- {
		child, parent := path[i], path[i-1]
		for j, ch := range parent.children {
			if ch == child {
				parent.children = append(parent.children[:j], parent.children[j+1:]...)
				break
			}
		}
		child.children = append(child.children, parent)
	}
	return path[len(path)-1]
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	fromTree := tr.FromPov(from)
	if fromTree == nil {
		return nil
	}

	path := fromTree.createPath(to, make([]*Tree, 0))
	p := make([]string, 0, len(path))
	for _, node := range path {
		p = append(p, node.value)
	}
	return p
}
