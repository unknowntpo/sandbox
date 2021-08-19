package main

import "fmt"

type LRUCache struct {
	capacity int
	index    map[int]*node
	head     *node
	tail     *node
}

type node struct {
	key, value int
	pre, next  *node
}

func newLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		index:    make(map[int]*node),
		head:     nil,
		tail:     nil,
	}
}

func (l *LRUCache) put(key, value int) {
	// if key exist, update value then return
	if n, ok := l.index[key]; ok {
		n.value = value
		return
	}

	n := &node{key: key, value: value}
	l.appendFront(n)
	l.addIndex(n)

	if !l.isFull() {
		return
	} else {
		evicted := l.evict()
		l.deleteIndex(evicted)
		return
	}
	return
}

// appendFront adds node n to the front of cache.
func (l *LRUCache) appendFront(n *node) {}

// addIndex adds reference to node n to index.
func (l *LRUCache) addIndex(n *node) {}

// deleteIndex delete the index of given node.
func (l *LRUCache) deleteIndex(n *node) {}

// ifFull return whether the cache is full or not.
// FIXME: not finished.
func (l *LRUCache) isFull() bool {
	return false
}

// evict evicts the tail node of cache, then return the evicted nodes' address.
// FIXME: not finished.
func (l *LRUCache) evict() *node {
	return nil
}

// show shows all elements in cache.
func (l *LRUCache) show() {
	for e := l.head; e != nil; e = e.next {
		fmt.Printf("k: %d, v: %d ->", e.key, e.value)
	}
}

func main() {
	/*
		LRUCache lRUCache = new LRUCache(2);
		lRUCache.put(1, 1); // cache is {1=1}
		lRUCache.put(2, 2); // cache is {1=1, 2=2}
		lRUCache.get(1);    // return 1
		lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
		lRUCache.get(2);    // returns -1 (not found)
		lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
		lRUCache.get(1);    // return -1 (not found)
		lRUCache.get(3);    // return 3
		lRUCache.get(4);    // return 4
	*/
	l := newLRUCache(5)
	l.show()
}
