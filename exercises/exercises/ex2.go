package exercices

/*
Is Unique: Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/

func HasUniqueCharWithDS(str string) bool {
	sMap := make(map[string]int, len(str))

	for i := 0; i < len(str); i++ {
		if _, ok := sMap[string(str[i])]; !ok {
			sMap[string(str[i])] = 0
		}
		sMap[string(str[i])]++
		if sMap[string(str[i])] > 1 {
			return false
		}
	}

	return true
}

func HasUniqueChar(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
