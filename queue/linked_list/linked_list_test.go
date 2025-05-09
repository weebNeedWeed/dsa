package linked_list

import (
	"testing"
)

func TestPushBackAndPopFront(t *testing.T) {
	list := NewLinkedList()
	data := []int{10, 20, 30, 40}
	for _, v := range data {
		list.PushBack(v)
	}

	for _, expected := range data {
		actual := list.PopFront()
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	}
}

func TestPopFrontEmptyList(t *testing.T) {
	list := NewLinkedList()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when popping from empty list, but did not panic")
		}
	}()
	list.PopFront() // should panic
}

func TestMultiplePushBack(t *testing.T) {
	list := NewLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	if list.size != 3 {
		t.Errorf("expected size 3, got %d", list.size)
	}
	if list.head.data != 1 {
		t.Errorf("expected head data 1, got %d", list.head.data)
	}
	if list.tail.data != 3 {
		t.Errorf("expected tail data 3, got %d", list.tail.data)
	}
}

func TestPushAndPopInterleaved(t *testing.T) {
	list := NewLinkedList()
	list.PushBack(100)
	if data := list.PopFront(); data != 100 {
		t.Errorf("expected 100, got %d", data)
	}

	list.PushBack(200)
	list.PushBack(300)
	if data := list.PopFront(); data != 200 {
		t.Errorf("expected 200, got %d", data)
	}

	list.PushBack(400)
	if data := list.PopFront(); data != 300 {
		t.Errorf("expected 300, got %d", data)
	}
	if data := list.PopFront(); data != 400 {
		t.Errorf("expected 400, got %d", data)
	}
}
