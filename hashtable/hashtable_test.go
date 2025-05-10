package main

import (
	"testing"
)

func TestNewHashTable(t *testing.T) {
	size := 10
	ht := NewHashTable(size)

	if ht.Size != size {
		t.Errorf("Expected size %d, got %d", size, ht.Size)
	}

	if len(ht.Inner) != size {
		t.Errorf("Expected inner array of length %d, got %d", size, len(ht.Inner))
	}

	for i, list := range ht.Inner {
		if list == nil {
			t.Errorf("LinkedList at index %d is nil", i)
		}
		if list.Size != 0 {
			t.Errorf("Expected empty list at index %d, got size %d", i, list.Size)
		}
	}
}

func TestHashFunction(t *testing.T) {
	size := 10
	ht := NewHashTable(size)

	testCases := []struct {
		key      int
		expected int
	}{
		{5, 5},
		{10, 0},
		{15, 5},
		{21, 1},
		{100, 0},
	}

	for _, tc := range testCases {
		result := ht.hash(tc.key)
		if result != tc.expected {
			t.Errorf("hash(%d) = %d; expected %d", tc.key, result, tc.expected)
		}
	}
}

func TestAddItem(t *testing.T) {
	ht := NewHashTable(10)

	// Add a simple item
	ht.Add(5, 50)

	// Check if item was added correctly
	idx := ht.hash(5)
	list := ht.Inner[idx]

	if list.Size != 1 {
		t.Errorf("Expected list size 1, got %d", list.Size)
	}

	if list.Head.Key.Key != 5 || list.Head.Key.Value != 50 {
		t.Errorf("Item not added correctly, got key=%d, value=%d", list.Head.Key.Key, list.Head.Key.Value)
	}
}

func TestCollision(t *testing.T) {
	ht := NewHashTable(10)

	// These will hash to the same slot
	ht.Add(5, 50)
	ht.Add(15, 150)

	// Check if both items exist in the same slot
	idx := ht.hash(5)
	list := ht.Inner[idx]

	if list.Size != 2 {
		t.Errorf("Expected list size 2, got %d", list.Size)
	}

	// Check if both items can be found
	current := list.Head
	foundFirst := false
	foundSecond := false

	list.Print()

	for current != nil {
		if current.Key.Key == 5 && current.Key.Value == 50 {
			foundFirst = true
		}
		if current.Key.Key == 15 && current.Key.Value == 150 {
			foundSecond = true
		}
		current = current.Next
	}

	if !foundFirst || !foundSecond {
		t.Errorf("Not all items found in collision list: first=%v, second=%v", foundFirst, foundSecond)
	}
}

func TestOverwriteExistingKey(t *testing.T) {
	ht := NewHashTable(10)

	// Add initial item
	ht.Add(5, 50)

	// Overwrite with new value
	ht.Add(5, 500)

	// Verify that value was updated and only one item exists
	idx := ht.hash(5)
	list := ht.Inner[idx]

	if list.Size != 1 {
		t.Errorf("Expected list size 1 after overwrite, got %d", list.Size)
	}

	if list.Head.Key.Key != 5 || list.Head.Key.Value != 500 {
		t.Errorf("Item not overwritten correctly, got key=%d, value=%d", list.Head.Key.Key, list.Head.Key.Value)
	}
}

func TestConsecutiveAdds(t *testing.T) {
	ht := NewHashTable(10)

	for i := 0; i < 20; i++ {
		ht.Add(i, i*10)
	}

	// Verify all items were added
	for i := 0; i < 20; i++ {
		idx := ht.hash(i)
		list := ht.Inner[idx]

		found := false
		current := list.Head
		for current != nil {
			if current.Key.Key == i && current.Key.Value == i*10 {
				found = true
				break
			}
			current = current.Next
		}

		if !found {
			t.Errorf("Item with key %d not found", i)
		}
	}
}
