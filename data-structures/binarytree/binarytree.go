package binarytree

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree
	value int
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{
		left:  nil,
		right: nil,
		value: value,
	}
}
func (tree *BinaryTree) Add(value int) {
	var pointer *BinaryTree

	if value > tree.value {
		pointer = tree.right
	} else {
		pointer = tree.left
	}

	if pointer == nil {
		pointer = &BinaryTree{
			left:  nil,
			right: nil,
			value: value,
		}
	} else {
		pointer.Add(value)
	}
}

// HasValue
// ToString
// Height
// Min
// Max
// Sum
