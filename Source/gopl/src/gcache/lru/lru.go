package lru

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

var KeyNotFoundError = errors.New("Key not found.")

type CacheNode struct {
	Key, Value interface{}
}

type LRUCache struct {
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
	mux      sync.RWMutex
}

var (
	CacheIsNULL = errors.New("Cache list is null")
)

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element, cap+1),
	}
}

func (lru *LRUCache) Size() int {
	return lru.dlist.Len()
}

func (lru *LRUCache) setValue(k, v interface{}) error {
	if lru == nil {
		return CacheIsNULL
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		pElement.Value.(*CacheNode).Value = v
		return nil
	}

	newElement := lru.dlist.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement

	if lru.dlist.Len() > lru.Capacity {
		lastElement := lru.dlist.Back()
		if lastElement != nil {
			cacheNode := lastElement.Value.(*CacheNode)
			delete(lru.cacheMap, cacheNode)
			lru.dlist.Remove(lastElement)
		}
	}
	return nil
}

// set a new key-value pair
func (lru *LRUCache) Set(key, value interface{}) error {
	lru.mux.Lock()
	defer lru.mux.Unlock()
	return lru.setValue(key, value)
}

func (lru *LRUCache) getValue(k interface{}) (v interface{}, err error) {
	if lru == nil {
		return v, CacheIsNULL
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		return pElement.Value.(*CacheNode).Value, nil
	}
	return v, KeyNotFoundError
}

// Get a value from cache pool using key if it exists.
func (lru *LRUCache) Get(key interface{}) (interface{}, error) {
	lru.mux.Lock()
	defer lru.mux.Unlock()
	return lru.getValue(key)
}

func (lru *LRUCache) remove(k interface{}) bool {
	if lru == nil {
		return false
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode)
		lru.dlist.Remove(pElement)
		return true
	}
	return false
}

// Removes the provided key from the cache.
func (lru *LRUCache) Remove(key interface{}) bool {
	lru.mux.Lock()
	defer lru.mux.Unlock()

	return lru.remove(key)
}

func (lru *LRUCache) Print() {
	if lru == nil {
		return
	}

	for e := lru.dlist.Front(); e != nil; e = e.Next() {
		cacheNode := e.Value.(*CacheNode)
		fmt.Println("Key:", cacheNode.Key, "Value:", cacheNode.Value)
	}
}
