package coalesced_chaining

type HtItem struct {
	Key, Next int
}

func NewHtItem(key, next int) *HtItem {
	return &HtItem{key, next}
}

type HashTable struct {
	Inner         []*HtItem
	Size          int
	LastAvailable int
}

func NewHashTable(size int) *HashTable {
	inner := make([]*HtItem, size)
	for i, _ := range inner {
		inner[i] = NewHtItem(-1, -1)
	}

	return &HashTable{
		Inner:         inner,
		Size:          size,
		LastAvailable: size - 1,
	}
}

func (h *HashTable) hash(key int) int {
	return key % h.Size
}

func (h *HashTable) Insert(key int) {
	if h.LastAvailable < 0 {
		return
	}

	idx := h.hash(key)

	slot := h.Inner[idx]
	if slot.Key == -1 {
		slot.Key = key
	} else {
		h.Inner[h.LastAvailable].Key = key

		for h.Inner[idx].Next != -1 {
			idx = h.Inner[idx].Next
		}

		h.Inner[idx].Next = h.LastAvailable

		for h.LastAvailable >= 0 && h.Inner[h.LastAvailable].Key != -1 {
			h.LastAvailable--
		}
	}
}
