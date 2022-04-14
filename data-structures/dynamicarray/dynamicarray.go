package dynamicarray

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	size          int
	lenght        int
	rateUpscaling int
	array         []int
}

func NewDynamicArray(size int, rateUpscaling int) *DynamicArray {
	return &DynamicArray{
		size:          size,
		lenght:        0,
		rateUpscaling: rateUpscaling,
		array:         make([]int, size),
	}
}

func (array *DynamicArray) Get(position int) (int, error) {
	if position > array.size || position < 0 {
		return 0, errors.New("IndexOutOfBound")
	}

	return array.array[position], nil
}

func (array *DynamicArray) Add(value int) {
	if array.lenght >= array.size {
		array.upscaling()
	}
	array.array[array.lenght] = value
	array.lenght++
}

func (array *DynamicArray) upscaling() {
	newSize := array.size * array.rateUpscaling
	newArray := make([]int, newSize)

	for i := 0; i < array.size; i++ {
		newArray[i] = array.array[i]
	}

	array.array = newArray
	array.size = newSize
}

func (array *DynamicArray) Print() {
	fmt.Printf(`{
		size   = %d,
		length = %d,
		rate   = %d,
		values = [
			%v
		]  
	}`, array.size, array.lenght, array.rateUpscaling, array.array)
}
