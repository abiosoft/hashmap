// Simple implementation of HashMap in GoLang
package hashmap

import (
	"fmt"
)

//default number of buckets
const BucketSize = 2 ^ 4

type HashMap struct {
	m          map[int]*list
	BucketSize int
}

// Create new Hashmap
func NewHashMap() *HashMap {
	return &HashMap{
		make(map[int]*list),
		BucketSize,
	}
}

// Put an item into the map
func (h *HashMap) Put(key string, value interface{}) {
	bucket := len(key) % h.BucketSize
	mapEntry := &mapEntry{key, nil, value}
	if _, ok := h.m[bucket]; !ok {
		h.m[bucket] = &list{}
	}
	h.m[bucket].Insert(mapEntry)
}

// Retrieve an item from the map using key
func (h *HashMap) Get(key string) interface{} {
	bucket := len(key) % h.BucketSize
	if list, ok := h.m[bucket]; ok {
		for l := list.head; l != nil; l = l.next {
			if l.key == key {
				return l.value
			}
		}
	}
	return nil
}

// Delete from the map using key
func (h *HashMap) Delete(key string) {
	bucket := len(key) % h.BucketSize
	if _, ok := h.m[bucket]; ok {
		h.m[bucket].Delete(key)
	}
}

// Retrieve the numbers of key value mappings in the map
func (h *HashMap) Size() int {
	size := 0
	for _, v := range h.m {
		size += v.Size()
	}
	return size
}

func (h *HashMap) String() string {
	str := fmt.Sprintf("Buckets: %d, ", h.Size)
	for i, list := range h.m {
		str += fmt.Sprint("[Bucket->", i, ":{", list, "}] ")
	}
	return str
}

type mapEntry struct {
	key   string
	next  *mapEntry
	value interface{}
}

// LinkedList
type list struct {
	head *mapEntry
	size int
}

func (l *list) Insert(mapEntry *mapEntry) {
	if l.head == nil {
		l.head = mapEntry
	} else {
		for n := &l.head; n != nil; {
			if (*n).key == mapEntry.key {
				(*n).value = mapEntry.value
				return
			}
			if (*n).next == nil {
				(*n).next = mapEntry
				break
			}
			n = &(*n).next
		}
	}
	l.size++
}

func (l *list) Delete(key string) {
	if l.head.key == key {
		l.head = l.head.next
		l.size--
		return
	}
	n := l.head
	var prev *mapEntry = n
	for n != nil {
		if n.key == key {
			prev.next = n.next
			break
		}
		prev = n
		n = n.next
	}
	l.size--
}

func (l *list) Size() int {
	return l.size
}

func (l *list) String() string {
	if l == nil {
		return "nil"
	}
	str := ""
	for mapEntry := l.head; mapEntry != nil; mapEntry = mapEntry.next {
		str += fmt.Sprint("[", mapEntry.key, ":", mapEntry.value, "]")
	}
	return str
}
