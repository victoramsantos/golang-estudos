package hashmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoramsantos/golang-estudos/data-structures/hashmap"
)

func TestHashmap(t *testing.T) {

	t.Run("adding values", func(t *testing.T) {
		hashmap := hashmap.NewHashMap(5)

		hashmap.Add("abc")
		hashmap.Add("acb")
		hashmap.Add("aaa")
		hashmap.Add("bbb")

		hashmap.Print()

		assert.True(t, hashmap.HasValue("abc"))
		assert.True(t, hashmap.HasValue("acb"))
		assert.True(t, hashmap.HasValue("aaa"))
		assert.True(t, hashmap.HasValue("bbb"))
	})
}
