package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	t.Run("should increase size by 2 and has exact 2 elements", func(t *testing.T) {
		da := newDynamicArray()

		da.push(1)
		da.push(2)

		if da.array[0] != 1 || da.array[1] != 2 {
			t.Errorf("got %d,%d; want %d,%d", da.array[0], da.array[1], 1, 2)
		}

		if da.size != 2 {
			t.Errorf("size = %d; size must be %d", da.size, 2)
		}
	})

	t.Run("should resize if size is equal to capacity", func(t *testing.T) {
		da := newDynamicArray()

		// resize once containing 17 elements
		for i := 1; i <= 17; i++ {
			da.push(i)
		}

		if da.capacity != 32 {
			t.Errorf("capacity = %d; capacity must be %d", da.capacity, 32)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("should panic if index >= size", func(t *testing.T) {
		da := newDynamicArray()
		da.push(1)

		assert.Panics(t, func() {
			da.insert(1, 2)
		}, "no pacic when index >= size")
	})

	t.Run("should resize if size is equal to capacity", func(t *testing.T) {
		da := newDynamicArray()

		for i := 1; i <= 16; i++ {
			da.push(i)
		}

		da.insert(1, 2)

		if da.capacity != 32 {
			t.Errorf("capacity = %d; capacity must be %d", da.capacity, 32)
		}
	})

	t.Run("should have 2 items and exactly 1,2", func(t *testing.T) {
		da := newDynamicArray()

		da.push(2)
		da.insert(0, 1)

		if da.size != 2 {
			t.Errorf("size = %d, size must be %d", da.size, 2)
		}

		if da.array[0] != 1 || da.array[1] != 2 {
			t.Errorf("got %d,%d; want %d,%d", da.array[0], da.array[1], 1, 2)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("should panic when no items in the array", func(t *testing.T) {
		da := newDynamicArray()

		assert.Panics(t, func() {
			da.pop()
		}, "no panic when no items in the array")
	})

	t.Run("should shrink the array when size = 1/4 capacity", func(t *testing.T) {
		da := newDynamicArray()

		for i := 1; i <= 5; i++ {
			da.push(i)
		}

		da.pop()

		if da.capacity != 4 {
			t.Error("not shrink after popping the 5th item")
		}
	})

	t.Run("should pop the last element", func(t *testing.T) {
		da := newDynamicArray()

		da.push(1)

		got := da.pop()
		want := 1

		if got != want {
			t.Errorf("got %d; want %d", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("should panic if index >= size", func(t *testing.T) {
		da := newDynamicArray()
		da.push(1)

		assert.Panics(t, func() {
			da.delete(1)
		}, "no pacic when index >= size")
	})

	t.Run("should shrink the array when size = 1/4 capacity", func(t *testing.T) {
		da := newDynamicArray()

		for i := 1; i <= 5; i++ {
			da.push(i)
		}

		da.delete(1)

		if da.capacity != 4 {
			t.Error("not shrink after popping the 5th item")
		}
	})

	t.Run("should delete successfully", func(t *testing.T) {
		da := newDynamicArray()

		da.push(1)

		da.delete(0)

		if da.size == 1 || len(da.array) == 1 {
			t.Error("array still has one element after deleting")
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("should return index if found", func(t *testing.T) {
		da := newDynamicArray()

		da.push(1)
		da.push(2)
		da.push(3)

		got := da.find(2)
		want := 1

		if got != want {
			t.Errorf("got index of %d; want %d", got, want)
		}
	})

	t.Run("should return -1 if not found", func(t *testing.T) {
		da := newDynamicArray()

		got := da.find(1)
		want := -1

		if got != want {
			t.Errorf("got index of %d; want %d", got, want)
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("should remove 1 item successfully", func(t *testing.T) {
		da := newDynamicArray()

		da.push(1)

		da.remove(1)

		if da.size == 1 || len(da.array) == 1 {
			t.Error("array still has 1 element of 1")
		}
	})

	t.Run("should remove 2 items successfully", func(t *testing.T) {
		da := newDynamicArray()

		da.push(1)
		da.push(1)

		da.remove(1)

		if da.size == 2 || len(da.array) == 2 {
			t.Error("array still has 2 element of 1")
		}
	})
}

func TestResize(t *testing.T) {
	t.Run("should grow the array", func(t *testing.T) {
		da := newDynamicArray()

		da.resize(32)

		if da.capacity != 32 && len(da.array) == 32 {
			t.Error("array wasn't resized to the capacity of 32 elements")
		}
	})

	t.Run("should shrink the array", func(t *testing.T) {
		da := newDynamicArray()

		for i := 1; i <= 16; i++ {
			da.push(i)
		}

		da.resize(4)

		if da.capacity != 4 && len(da.array) == 4 {
			t.Error("array wasn't resized to the capacity of 4 elements")
		}

		if da.size != 4 {
			t.Error("the size of the array is not shrinked to 4")
		}
	})
}
