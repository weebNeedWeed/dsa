package main

type HtItem struct {
	Key, Value int
}

func NewHtItem(key, value int) HtItem {
	return HtItem{key, value}
}

type HashTable struct {
	Inner []*LinkedList
	Size  int
}

func NewHashTable(size int) *HashTable {
	inner := make([]*LinkedList, size)
	for i, _ := range inner {
		inner[i] = NewLinkedList()
	}

	return &HashTable{
		Inner: inner,
		Size:  size,
	}
}

func (h *HashTable) hash(key int) int {
	return key % h.Size
}

func (h *HashTable) Add(key, value int) {
	hashRes := h.hash(key)
	slot := h.Inner[hashRes]

	itemToAdd := NewHtItem(key, value)

	// no item -> should add
	if slot.Size == 0 {
		slot.PushFront(itemToAdd)
	} else {
		// check if key-value pair existed
		if slot.SearchByKeyOfHtItem(key) { // existed
			slot.DeleteByHtItemKey(key)
		}
		slot.PushFront(itemToAdd)
	}
}
