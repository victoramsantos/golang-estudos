package sortmethod_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoramsantos/golang-estudos/data-structures/sortmethod"
)

func TestSort(t *testing.T) {

	t.Run("bubblesort - sorted", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		sortmethod.BubbleSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("bubblesort - unsorted", func(t *testing.T) {
		array := []int{6, 4, 5, 2, 1, 3, 7, 8}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		sortmethod.BubbleSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("bubblesort - desc sorted", func(t *testing.T) {
		array := []int{8, 7, 6, 5, 4, 3, 2, 1}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		sortmethod.BubbleSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("quicksort - sorted", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		sortmethod.QuickSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("quicksort - unsorted", func(t *testing.T) {
		array := []int{6, 4, 5, 2, 1, 3, 7, 8}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		sortmethod.QuickSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("quicksort - desc sorted", func(t *testing.T) {
		array := []int{8, 7, 6, 5, 4, 3, 2, 1}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		sortmethod.QuickSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("mergesort - sorted", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		array = sortmethod.MergeSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("mergesort - unsorted", func(t *testing.T) {
		array := []int{6, 4, 5, 2, 1, 3, 7, 8}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		array = sortmethod.MergeSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})

	t.Run("mergesort - desc sorted", func(t *testing.T) {
		array := []int{8, 7, 6, 5, 4, 3, 2, 1}
		sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

		array = sortmethod.MergeSort(array)

		for i := 0; i < len(array); i++ {
			assert.Equal(t, array[i], sortedArray[i])
		}
	})
}
