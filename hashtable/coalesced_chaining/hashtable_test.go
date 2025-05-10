package coalesced_chaining

import (
	"testing"
)

func TestNewHashTable(t *testing.T) {
	size := 10
	ht := NewHashTable(size)

	if ht.Size != size {
		t.Errorf("Expected size to be %d, got %d", size, ht.Size)
	}

	if ht.LastAvailable != size-1 {
		t.Errorf("Expected LastAvailable to be %d, got %d", size-1, ht.LastAvailable)
	}

	for i, item := range ht.Inner {
		if item.Key != -1 {
			t.Errorf("Expected Key at position %d to be -1, got %d", i, item.Key)
		}
		if item.Next != -1 {
			t.Errorf("Expected Next at position %d to be -1, got %d", i, item.Next)
		}
	}
}

func TestHashFunction(t *testing.T) {
	ht := NewHashTable(10)

	testCases := []struct {
		key      int
		expected int
	}{
		{5, 5},
		{15, 5},
		{25, 5},
		{11, 1},
		{0, 0},
	}

	for _, tc := range testCases {
		result := ht.hash(tc.key)
		if result != tc.expected {
			t.Errorf("hash(%d) expected %d, got %d", tc.key, tc.expected, result)
		}
	}
}

func TestInsert(t *testing.T) {
	ht := NewHashTable(5)

	// Insert key that maps directly to its hash
	ht.Insert(3)
	if ht.Inner[3].Key != 3 {
		t.Errorf("Expected key 3 at position 3, got %d", ht.Inner[3].Key)
	}

	// Insert another key with the same hash (collision)
	ht.Insert(8) // 8 % 5 = 3, so this will collide with the previous key
	if ht.Inner[4].Key != 8 {
		t.Errorf("Expected key 8 to be placed at position 4 (LastAvailable), got key %d", ht.Inner[4].Key)
	}

	if ht.Inner[3].Next != 4 {
		t.Errorf("Expected next pointer from position 3 to point to 4, got %d", ht.Inner[3].Next)
	}

	// Test that LastAvailable is updated correctly
	if ht.LastAvailable != 2 {
		t.Errorf("Expected LastAvailable to be updated to 3, got %d", ht.LastAvailable)
	}
}

func TestInsertChaining(t *testing.T) {
	ht := NewHashTable(5)

	// Create a long chain of collisions at position 0
	ht.Insert(0)  // hash = 0
	ht.Insert(5)  // hash = 0
	ht.Insert(10) // hash = 0

	// Check direct placement
	if ht.Inner[0].Key != 0 {
		t.Errorf("Expected key 0 at position 0, got %d", ht.Inner[0].Key)
	}

	// Check first collision resolved to position 4
	if ht.Inner[4].Key != 5 {
		t.Errorf("Expected key 5 at position 4, got %d", ht.Inner[4].Key)
	}

	// Check second collision resolved to position 3
	if ht.Inner[3].Key != 10 {
		t.Errorf("Expected key 10 at position 3, got %d", ht.Inner[3].Key)
	}

	// Follow the chain
	if ht.Inner[0].Next != 4 {
		t.Errorf("Expected next pointer from position 0 to point to 4, got %d", ht.Inner[0].Next)
	}

	if ht.Inner[4].Next != 3 {
		t.Errorf("Expected next pointer from position 4 to point to 3, got %d", ht.Inner[4].Next)
	}
}

func TestFullHashTable(t *testing.T) {
	ht := NewHashTable(3)

	// Fill the table
	ht.Insert(0)
	ht.Insert(1)
	ht.Insert(2)

	// This should silently not insert as the table is full
	ht.Insert(3)

	// Verify last available is -1 (full table)
	if ht.LastAvailable != -1 {
		t.Errorf("Expected LastAvailable to be -1 for full table, got %d", ht.LastAvailable)
	}
}
