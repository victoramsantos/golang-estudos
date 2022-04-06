package binarytree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoramsantos/golang-estudos/data-structures/binarytree"
)

func TestBinaryTree(t *testing.T) {

	t.Run("printing values", func(t *testing.T) {
		values := [6]int{5, 3, 6, 1, 2, 4}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}
		tree.ToString()
	})

	t.Run("adding values", func(t *testing.T) {
		values := [6]int{5, 3, 6, 1, 2, 4}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, len(values), tree.CountMembers())
	})

	t.Run("minimum value", func(t *testing.T) {
		values := [6]int{5, 3, 6, 1, 2, 4}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, len(values), tree.CountMembers())
		assert.Equal(t, 1, tree.Min())
	})

	t.Run("maximum value", func(t *testing.T) {
		values := [6]int{5, 3, 6, 1, 2, 4}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, len(values), tree.CountMembers())
		assert.Equal(t, 6, tree.Max())
	})

	t.Run("has value", func(t *testing.T) {
		values := [6]int{5, 3, 7, 1, 2, 4}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, len(values), tree.CountMembers())
		assert.True(t, tree.HasValue(1))
		assert.True(t, tree.HasValue(2))
		assert.True(t, tree.HasValue(3))
		assert.True(t, tree.HasValue(4))
		assert.True(t, tree.HasValue(5))
		assert.True(t, tree.HasValue(7))

		assert.False(t, tree.HasValue(0))
		assert.False(t, tree.HasValue(6))
		assert.False(t, tree.HasValue(8))
	})

	t.Run("height - only root", func(t *testing.T) {
		values := []int{5}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, 1, tree.Height())
	})

	t.Run("height - unsorted", func(t *testing.T) {
		values := []int{10, 4, 5, 3, 2, 1, 0, -1}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, 7, tree.Height())
	})

	t.Run("height - ordered", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6, 7, 8}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		assert.Equal(t, 8, tree.Height())
	})

	t.Run("flattening tree", func(t *testing.T) {
		values := []int{5, 3, 7, 1, 2, 4}
		tree := binarytree.NewBinaryTree(values[0])

		for i := 1; i < len(values); i++ {
			tree.Add(values[i])
		}

		array := tree.Flatten()

		assert.Equal(t, tree.CountMembers(), len(array))
		assert.Contains(t, array, 5)
		assert.Contains(t, array, 3)
		assert.Contains(t, array, 7)
		assert.Contains(t, array, 1)
		assert.Contains(t, array, 2)
		assert.Contains(t, array, 4)
	})

}
