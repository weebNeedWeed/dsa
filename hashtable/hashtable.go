package main

// type HashTable struct {
// 	Inner []*LinkedList
// 	Size  int
// }

// func NewHashTable(size int) *HashTable {
// 	inner := make([]*LinkedList, size)
// 	for i, _ := range inner {
// 		inner[i] = NewLinkedList(nil)
// 	}

// 	return &HashTable{
// 		Inner: inner,
// 		Size:  size,
// 	}
// }

// func (h *HashTable) hash(key int) int {
// 	return key % h.Size
// }

// func (h *HashTable) Add(key, value int) {
// 	hashRes := h.hash(key)
// 	slot := h.Inner[hashRes]

// 	if slot.Search(value) == true {

// 	}
// }
