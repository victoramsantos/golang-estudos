package binarytree

import "fmt"

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
	var pointer **BinaryTree

	if value > tree.value {
		pointer = &tree.right
	} else {
		pointer = &tree.left
	}

	if *pointer == nil {
		*pointer = &BinaryTree{
			left:  nil,
			right: nil,
			value: value,
		}
	} else {
		(*pointer).Add(value)
	}
}

func (tree *BinaryTree) ToString() {
	tree._toString("")
}

func (tree *BinaryTree) _toString(line string) {
	fmt.Printf("%d\n", tree.value)

	if tree.left != nil {
		fmt.Printf("%s|__", line)
		tree.left._toString(fmt.Sprintf("  %s", line))
	}

	if tree.right != nil {
		fmt.Printf("%s|__", line)
		tree.right._toString(fmt.Sprintf("  %s", line))
	}
}

func (tree *BinaryTree) CountMembers() int {
	sum := 0
	if tree.left != nil {
		sum += tree.left.CountMembers()
	}
	if tree.right != nil {
		sum += tree.right.CountMembers()
	}
	return sum + 1
}

func (tree *BinaryTree) Min() int {
	min := tree.value
	tree._min(&min)
	return min
}

func (tree *BinaryTree) _min(value *int) {
	if tree.value < *value {
		*value = tree.value
	}
	if tree.left != nil {
		tree.left._min(value)
	}
	if tree.right != nil {
		tree.right._min(value)
	}
}

func (tree *BinaryTree) Max() int {
	max := tree.value
	tree._max(&max)
	return max
}

func (tree *BinaryTree) _max(value *int) {
	if tree.value > *value {
		*value = tree.value
	}
	if tree.left != nil {
		tree.left._max(value)
	}
	if tree.right != nil {
		tree.right._max(value)
	}
}

func (tree *BinaryTree) HasValue(value int) bool {
	if tree.value == value {
		return true
	} else {
		if value > tree.value {
			if tree.right != nil {
				return tree.right.HasValue(value)
			}
		} else {
			if tree.left != nil {
				return tree.left.HasValue(value)
			}
		}
	}
	return false
}

// HasValue
// Height
