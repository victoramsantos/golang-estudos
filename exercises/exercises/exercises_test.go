package exercices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	exercices "github.com/victoramsantos/golang-estudos/exercises/exercises"
)

func Test_ex1(t *testing.T) {
	str := "abc"
	expectedPerms := []string{"cab", "acb", "abc", "cba", "bca", "bac"}
	expected := make(map[string]struct{}, len(expectedPerms))
	for _, s := range expectedPerms {
		expected[s] = struct{}{}
	}

	t.Run("permute", func(t *testing.T) {
		perms := exercices.Permute(str)

		assert.Equal(t, len(expectedPerms), len(perms))

		for _, perm := range perms {
			_, has := expected[perm]
			assert.True(t, has)
		}
	})
}

func Test_ex2(t *testing.T) {
	uniqueChars1 := "abcdefghij"
	uniqueChars2 := "a"
	uniqueChars3 := ""
	nonUniqueChars1 := "abcdeefghij"

	t.Run("has unique", func(t *testing.T) {
		assert.True(t, exercices.HasUniqueCharWithDS(uniqueChars1))
		assert.True(t, exercices.HasUniqueCharWithDS(uniqueChars2))
		assert.True(t, exercices.HasUniqueCharWithDS(uniqueChars3))

		assert.True(t, exercices.HasUniqueChar(uniqueChars1))
		assert.True(t, exercices.HasUniqueChar(uniqueChars2))
		assert.True(t, exercices.HasUniqueChar(uniqueChars3))
	})

	t.Run("no has unique", func(t *testing.T) {
		assert.False(t, exercices.HasUniqueCharWithDS(nonUniqueChars1))
		assert.False(t, exercices.HasUniqueChar(nonUniqueChars1))
	})
}

func Test_ex3(t *testing.T) {
	strPair1 := []string{"abc", "cba"}
	strPair2 := []string{"aaab", "aaba"}
	strPair3 := []string{"a", "a"}
	strPair4 := []string{"ab", "ba"}
	strPair5 := []string{"", ""}
	strPair6 := []string{"a", "ab"}
	strPair7 := []string{"b", "ba"}
	strPair8 := []string{"a", ""}
	strPair9 := []string{"", "a"}

	t.Run("has unique", func(t *testing.T) {
		assert.True(t, exercices.IsPermsOf(strPair1[0], strPair1[1]))
		assert.True(t, exercices.IsPermsOf(strPair2[0], strPair2[1]))
		assert.True(t, exercices.IsPermsOf(strPair3[0], strPair3[1]))
		assert.True(t, exercices.IsPermsOf(strPair4[0], strPair4[1]))
		assert.True(t, exercices.IsPermsOf(strPair5[0], strPair5[1]))

		assert.False(t, exercices.IsPermsOf(strPair6[0], strPair6[1]))
		assert.False(t, exercices.IsPermsOf(strPair7[0], strPair7[1]))
		assert.False(t, exercices.IsPermsOf(strPair8[0], strPair8[1]))
		assert.False(t, exercices.IsPermsOf(strPair9[0], strPair9[1]))
	})

}

func Test_ex4(t *testing.T) {
	str1 := "Mr John Smith    "
	str2 := "a b c        "

	t.Run("URLify", func(t *testing.T) {
		assert.Equal(t, "Mr%20John%20Smith", exercices.URLify(str1, 13))
		assert.Equal(t, "a%20b%20c", exercices.URLify(str2, 5))
	})

}
