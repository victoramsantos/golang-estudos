package sortmethod

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
}
