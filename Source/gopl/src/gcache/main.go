package main

import (
	"fmt"

	"gcache/lru"
)

func main() {
	cache := lru.NewLRUCache(5)
	cache.Set(10, "Value 10")
	cache.Set(20, "Value 20")
	cache.Set(30, "Value 30")
	cache.Set(40, "Value 40")
	cache.Set(50, "Value 50")
	cache.Set(10, "Value 60")

	fmt.Println("LRU cache size: ", cache.Size())
	cache.Print()

	//	if v, ret, _ := cache.Get(30); ret {
	//		fmt.Println("Get(30) : ", v)
	//	}
	//
	//	if cache.Remove(30) {
	//		fmt.Println("Remove(30) : true")
	//	} else {
	//		fmt.Println("Remove(30) : false")
	//	}
	//	fmt.Println("LRU Size:", cache.Size())

	str := "abcd"
	var tem = []byte(str)
	fmt.Println(tem)
}
