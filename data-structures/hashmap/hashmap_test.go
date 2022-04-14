package hashmap_test

import (
	"testing"

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
	})
}
