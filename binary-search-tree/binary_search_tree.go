package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {

	node := bst
	newNode := NewBst(i)

	for {
		if i <= node.data {
			if node.left != nil {
				node = node.left
			} else {
				node.left = newNode
				break
			}
		} else {
			if node.right != nil {
				node = node.right
			} else {
				node.right = newNode
				break
			}
		}
	}
}

func sortTree(node *BinarySearchTree, s *[]int) {
	if node != nil {
		sortTree(node.left, s)
		*s = append(*s, node.data)
		sortTree(node.right, s)
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	s := []int{}
	sortTree(bst, &s)

	return s
}
