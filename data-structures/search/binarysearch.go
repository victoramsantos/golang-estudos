package search

func BinarySearch(array []int, value int) bool {
	return _binarySearch(array, value, 0, len(array))
}

func _binarySearch(array []int, value int, posBegin int, posEnd int) bool {
	if posBegin < 0 || posEnd < 0 || posBegin == len(array) {
		return false
	}

	mid := (posEnd + posBegin) / 2

	if value == array[mid] {
		return true
	}
	if value > array[mid] {
		return _binarySearch(array, value, mid+1, posEnd)
	} else {
		return _binarySearch(array, value, posBegin, mid-1)
	}
}
