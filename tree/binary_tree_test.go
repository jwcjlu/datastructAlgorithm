package tree

import (
	"fmt"
	"testing"
)

type Key int

func (k Key) Compare(j Key) int {
	if k == j {
		return 0
	}
	if k < j {
		return 1
	}
	return -1
}
func ofKey(k int) Key {
	return Key(k)
}
func Test_Bst(t *testing.T) {
	bst := BinaryTree{Compare: func(i interface{}, i2 interface{}) int {
		v1, ok := i.(int)
		if !ok {
			return -1
		}
		v2, ok := i2.(int)
		if !ok {
			return -1
		}
		if v1 == v2 {
			return 0
		}
		if v1 > v2 {
			return 1
		} else {
			return -1
		}

	}}
	bst.Put(1, 1)
	bst.Put(2, 2)
	bst.Put(7, 7)
	bst.Put(3, 3)
	bst.Put(5, 5)
	bst.Put(8, 8)
	bst.Put(9, 9)
	fmt.Println(bst)
}
