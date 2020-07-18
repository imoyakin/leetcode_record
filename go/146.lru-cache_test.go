package leetcode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	//.Put(2, 2)
	cache.Get(2)    // 返回  1
	cache.Put(3, 2) // 该操作会使得关键字 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	// cache.Put(4, 4) // 该操作会使得关键字 1 作废
	// cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3) // 返回  3
	// cache.Get(4)    // 返回  4
	// fmt.Println(cache)
}
