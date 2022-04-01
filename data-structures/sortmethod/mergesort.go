package sortmethod

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(array); i++ {
		if i < len(array)/2 {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
