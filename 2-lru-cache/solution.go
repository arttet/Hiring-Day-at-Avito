// Package problem contains one coding problem.
package problem

import (
	"container/list"
	"strings"
)

const capacity = 5

// Cache is a LRU cache.
type Cache struct {
	cap int
	l   *list.List
	m   map[string]*list.Element
}

func NewCache(capacity int) Cache {
	return Cache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[string]*list.Element, capacity),
	}
}

func (cache *Cache) String() string {
	var buf strings.Builder
	buf.Grow(2*cache.cap - 1)

	for e := cache.l.Back(); e != nil; {
		str := e.Value.(*list.Element).Value.(string)
		buf.WriteString(str)

		e = e.Prev()
		if e != nil {
			buf.WriteByte('-')
		}
	}

	return buf.String()
}

func (cache *Cache) Put(key string) {
	if node, ok := cache.m[key]; ok {
		cache.l.MoveToFront(node)
		node.Value.(*list.Element).Value = key
	} else {
		if cache.l.Len() == cache.cap {
			idx := cache.l.Back().Value.(*list.Element).Value.(string)
			delete(cache.m, idx)
			cache.l.Remove(cache.l.Back())
		}
		node := &list.Element{
			Value: key,
		}

		ptr := cache.l.PushFront(node)
		cache.m[key] = ptr
	}
}

func LRUCache(strArr []string) string {
	cache := NewCache(capacity)

	for _, str := range strArr {
		cache.Put(str)
	}

	result := cache.String()
	return result
}
