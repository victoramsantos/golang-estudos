package sortmethod

func QuickSort(array []int) {
	_quickSort(array, 0, len(array)-1)
}

func _quickSort(array []int, posBeg int, posEnd int) {
	var i, j, pivo, aux int

	i = posBeg
	j = posEnd

	pivo = array[(posBeg+posEnd)/2]

	for i <= j {
		for array[i] < pivo {
			i += 1
		}
		for array[j] > pivo {
			j -= 1
		}

		if i <= j {
			aux = array[i]
			array[i] = array[j]
			array[j] = aux
			i += 1
			j -= 1
		}
	}
	if posBeg < j {
		_quickSort(array, posBeg, j)
	}
	if i < posEnd {
		_quickSort(array, i, posEnd)
	}
}
