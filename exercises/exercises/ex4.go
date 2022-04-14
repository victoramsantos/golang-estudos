package exercices

/*
URLify: Write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the additional characters,
and that you are given the "true" length of the string.
EXAMPLE
Input: "Mr John Smith ", 13
Output: "Mr%20John%20Smith"
*/

/*
	spacePositions= [2,7]
	spaceLength=2
*/

func URLify(str string, length int) string {
	spacePositions := make(map[int]struct{}, length)
	spaceLength := 0
	s := ""

	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			spacePositions[i] = struct{}{}
			spaceLength++
		}
	}

	for i := length - 1; i >= 0; i-- {
		if _, ok := spacePositions[i]; ok {
			s += "02%"
		} else {
			s += string(str[i])
		}
	}

	last := 2*spaceLength + length
	return reverse(s, last)
}

func reverse(s string, length int) string {
	reverse := ""
	for i := length - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
