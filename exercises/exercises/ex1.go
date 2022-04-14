package exercices

/*
	Design an algorithm to print all permutations of a string.
	For simplicity, assume all characters are unique.
*/

func Permute(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}
	letter := str[len(str)-1:]
	perms := Permute(str[0 : len(str)-1])
	outPerms := []string{}
	for i := 0; i < len(perms); i++ {
		perm := perms[i]
		for j := 0; j <= len(perm); j++ {
			outPerms = append(outPerms, perm[0:j]+letter+perm[j:])
		}
	}
	return outPerms
}
