package hashmap

import (
	"fmt"
	"strings"

	"github.com/victoramsantos/golang-estudos/data-structures/linkedlist"
)

type Hashmap struct {
	arrayMap []*linkedlist.LinkedList
	size     int
}

func NewHashMap(size int) *Hashmap {
	array := make([]*linkedlist.LinkedList, size)

	for i := 0; i < size; i++ {
		array[i] = nil
	}

	return &Hashmap{
		size:     size,
		arrayMap: array,
	}
}

func (hashmap *Hashmap) Add(str string) {
	hash := calcHash(str)
	fmt.Printf("%s:%d\n", str, hash)
	position := hash % hashmap.size

	if hashmap.arrayMap[position] == nil {
		hashmap.arrayMap[position] = linkedlist.NewLinkedList(hash)
	} else {
		hashmap.arrayMap[position].Add(hash)
	}
}

func calcHash(str string) int {
	hash := 0
	for _, ch := range strings.ToLower(str) {
		hash += int(ch)
	}
	return hash
}

func (hashmap *Hashmap) Print() {
	maps := ""
	for i := 0; i < hashmap.size; i++ {
		maps += fmt.Sprintf("[%d] -> ", i)
		if hashmap.arrayMap[i] == nil {
			maps += "nil"
		} else {
			maps += hashmap.arrayMap[i].ToString()
		}
		maps += "\n"
	}

	fmt.Printf(`hashmap {
		size: %d,
		array: [
%s
		]
	}`, hashmap.size, maps)
}
