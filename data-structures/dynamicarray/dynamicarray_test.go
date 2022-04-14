package dynamicarray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoramsantos/golang-estudos/data-structures/dynamicarray"
)

func TestDynamicArray(t *testing.T) {

	t.Run("adding values", func(t *testing.T) {
		size := 5
		array := dynamicarray.NewDynamicArray(size, 2)

		for i := 0; i < size; i++ {
			array.Add(i)
		}

		array.Print()

		for i := 0; i < size; i++ {
			v, err := array.Get(i)
			assert.Nil(t, err)
			assert.Equal(t, i, v)
		}
	})

	t.Run("doubling size", func(t *testing.T) {
		size := 5
		array := dynamicarray.NewDynamicArray(size, 2)

		for i := 0; i < size*2; i++ {
			array.Add(i)
		}

		array.Print()

		for i := 0; i < size; i++ {
			v, err := array.Get(i)
			assert.Nil(t, err)
			assert.Equal(t, i, v)
		}
	})
}
