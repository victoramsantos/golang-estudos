package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoramsantos/golang-estudos/data-structures/linkedlist"
)

func TestLinkedList(t *testing.T) {

	t.Run("printing values", func(t *testing.T) {
		appends := 5
		list := linkedlist.NewLinkedList(0)

		for i := 1; i < appends; i++ {
			list.Add(i)
		}
		fmt.Println(list.ToString())
	})

	t.Run("length", func(t *testing.T) {
		appends := 5
		list := linkedlist.NewLinkedList(0)

		for i := 1; i < appends; i++ {
			list.Add(i)
		}

		assert.Equal(t, appends, list.Length())
	})

	t.Run("get at", func(t *testing.T) {
		appends := 5
		list := linkedlist.NewLinkedList(0)

		for i := 1; i < appends; i++ {
			list.Add(i)
		}

		value, err := list.GetValueAt(0)
		assert.Nil(t, err)
		assert.Equal(t, 0, value)

		value, err = list.GetValueAt(4)
		assert.Nil(t, err)
		assert.Equal(t, 4, value)

		_, err = list.GetValueAt(11)
		assert.NotNil(t, err)

		_, err = list.GetValueAt(-1)
		assert.NotNil(t, err)
	})

	t.Run("add at", func(t *testing.T) {
		appends := 5
		list := linkedlist.NewLinkedList(0)

		for i := 1; i < appends; i++ {
			list.Add(i)
		}

		err := list.AddAt(-1, 0)
		assert.Nil(t, err)
		value, err := list.GetValueAt(0)
		assert.Nil(t, err)
		assert.Equal(t, -1, value)
		assert.Equal(t, appends+1, list.Length())

		err = list.AddAt(-1, -1)
		assert.NotNil(t, err)
	})

	t.Run("new linkedlist", func(t *testing.T) {
		list := linkedlist.NewLinkedList(0)

		value, err := list.GetValueAt(0)
		assert.Nil(t, err)
		assert.Equal(t, 0, value)
		assert.Equal(t, 1, list.Length())
	})
}
