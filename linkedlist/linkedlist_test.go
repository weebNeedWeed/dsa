package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArray(t *testing.T) {
	t.Run("should convert linkedlist to array", func(t *testing.T) {
		l := newLinkedList(nil)

		n1 := newNode(1, nil)
		n2 := newNode(2, n1)

		l.head = n2

		got := l.toArray()
		want := []int{2, 1}

		if got[0] != 2 || got[1] != 1 || len(got) != 2 {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestPushFront(t *testing.T) {
	t.Run("should push at the front", func(t *testing.T) {
		l := newLinkedList(nil)

		l.pushFront(1)
		l.pushFront(2)

		got := l.toArray()
		want := []int{2, 1}

		if got[0] != 2 || got[1] != 1 || len(got) != 2 || l.size != 2 {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestValueAt(t *testing.T) {
	l := newLinkedList(nil)
	l.pushFront(1)
	l.pushFront(2)

	t.Run("should panic when index is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			l.valueAt(-1)
		}, "no panic when passing negative index")
	})

	t.Run("should return the head's data", func(t *testing.T) {
		got := l.valueAt(0)
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should panic when index is out of the list's length", func(t *testing.T) {
		assert.Panics(t, func() {
			l.valueAt(2)
		}, "no panic when passing out-of-length index")
	})

	t.Run("should return the second value", func(t *testing.T) {
		got := l.valueAt(1)
		want := 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestPushBack(t *testing.T) {
	t.Run("should set the head if there is no head", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushBack(1)

		if l.head == nil || l.head.data != 1 || l.size != 1 {
			t.Error("unable to push back in the list")
		}
	})

	t.Run("should push back 2 items", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushBack(1)
		l.pushBack(2)

		got := l.toArray()
		if got[0] != 1 || got[1] != 2 || l.size != 2 {
			t.Errorf("unable to push back 2 items in the list; got %v, want %d, %d", got, 1, 2)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("should panic when index is out of bound", func(t *testing.T) {
		l := newLinkedList(nil)

		assert.Panics(t, func() {
			l.insert(1, 2)
		}, "no panic when index is out of bound")
	})

	t.Run("should panic when no head was set", func(t *testing.T) {
		l := newLinkedList(nil)

		assert.Panics(t, func() {
			l.insert(0, 2)
		}, "no panic when no head was set")
	})

	t.Run("should insert new node at index 0", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushBack(1)
		l.insert(0, 2)

		got := l.toArray()
		if got[0] != 2 || got[1] != 1 || l.size != 2 {
			t.Errorf("unable to insert into the list; got %v, want %d, %d", got, 1, 2)
		}
	})

	t.Run("should insert new node at index 1", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushBack(1)
		l.pushBack(2)
		l.pushBack(3)
		l.insert(1, 4)

		got := l.toArray()
		want := []int{1, 4, 2, 3}
		if len(got) != len(want) {
			t.Errorf("unable to insert into the list; got %d items, want %d tiems", len(got), len(want))
		}
		for index, item := range got {
			if item != want[index] {
				t.Errorf("unable to insert into the list; got %v, want %v", got, want)
			}
		}
	})
	t.Run("should insert new node at the last index", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushBack(3)
		l.pushBack(4)
		l.insert(1, 5)

		got := l.toArray()
		want := []int{3, 5, 4}
		if len(got) != len(want) {
			t.Errorf("unable to insert into the list; got %d items, want %d tiems", len(got), len(want))
		}
		for index, item := range got {
			if item != want[index] {
				t.Errorf("unable to insert into the list; got %v, want %v", got, want)
			}
		}
	})
}

func TestPopFirst(t *testing.T) {
	t.Run("should panic if list is empty", func(t *testing.T) {
		l := newLinkedList(nil)

		assert.Panics(t, func() {
			l.popFirst()
		}, "no panic when head is empty")
	})

	t.Run("should pop first item succesfully", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushFront(1)

		got := l.popFirst()
		want := 1

		if got != want || l.size != 0 {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestPopBack(t *testing.T) {
	t.Run("should panic if list is empty", func(t *testing.T) {
		l := newLinkedList(nil)

		assert.Panics(t, func() {
			l.popBack()
		}, "no panic when head is empty")
	})

	t.Run("should pop last item succesfully (list has 1 items)", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushFront(1)

		got := l.popBack()
		want := 1

		if got != want || l.size != 0 {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should pop last item succesfully (list has 2 items)", func(t *testing.T) {
		l := newLinkedList(nil)
		l.pushFront(1)
		l.pushFront(2)

		got := l.popBack()
		want := 1

		if got != want || l.size != 1 {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
