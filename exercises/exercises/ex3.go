package exercices

/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
*/

/*
assumptions:
- there are no spaces as difference
- we'll ignore case
*/

func IsPermsOf(str1 string, str2 string) bool {
	compoundMap := make(map[string]int, len(str1))

	for _, ch := range str1 {
		if _, ok := compoundMap[string(ch)]; !ok {
			compoundMap[string(ch)] = 0
		}
		compoundMap[string(ch)]++
	}

	for _, ch := range str2 {
		if _, ok := compoundMap[string(ch)]; !ok {
			return false
		}
		compoundMap[string(ch)]++
	}

	for _, v := range compoundMap {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
