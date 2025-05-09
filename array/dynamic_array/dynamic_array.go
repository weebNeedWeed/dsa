package main

import (
	"log"
)

const DEFAULT_CAPACITY = 16

type dynamicArray struct {
	array    []int
	capacity uint
	size     uint
}

func newDynamicArray() *dynamicArray {
	return &dynamicArray{
		array:    make([]int, DEFAULT_CAPACITY, DEFAULT_CAPACITY),
		capacity: DEFAULT_CAPACITY,
		size:     0,
	}
}

func newDynamicArrayWithCapacity(capacity uint) *dynamicArray {
	return &dynamicArray{
		array:    make([]int, capacity, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (d *dynamicArray) getSize() uint {
	return d.size
}

func (d *dynamicArray) getCapacity() uint {
	return d.capacity
}

func (d *dynamicArray) isEmpty() bool {
	return d.size == 0
}

func (d *dynamicArray) at(index uint) int {
	if index >= d.size {
		log.Fatalf("given index, %d, is out of bound", index)
	}

	return d.array[index]
}

func (d *dynamicArray) push(item int) {
	if d.size == d.capacity {
		d.resize(d.capacity * 2)
	}

	d.array[d.size] = item
	d.size++
}

func (d *dynamicArray) insert(index uint, item int) {
	if index >= d.size {
		log.Panicf("given index, %d, is out of bound", index)
	}

	if d.size == d.capacity {
		d.resize(d.capacity * 2)
	}

	for i := d.size; i > index; i-- {
		d.array[i] = d.array[i-1]
	}

	d.size++
	d.array[index] = item
}

func (d *dynamicArray) prepend(item int) {
	d.insert(0, item)
}

func (d *dynamicArray) pop() int {
	if d.size == 0 {
		log.Panic("there is no item in the array")
	}

	result := d.array[d.size-1]
	d.size--

	if d.size*4 == d.capacity { // shrink if size = 1/4 capacity
		d.resize(d.capacity / 4)
	}

	return result
}

func (d *dynamicArray) delete(index uint) {
	if index >= d.size {
		log.Panicf("given index, %d, is out of bound", index)
	}

	for i := int(index); i < int(d.size)-2; i++ {
		d.array[i] = d.array[i+1]
	}

	d.size--

	if d.size*4 == d.capacity { // shrink if size = 1/4 capacity
		d.resize(d.capacity / 4)
	}
}

func (d *dynamicArray) find(item int) int {
	index := -1

	var i uint
	for i = 0; i < d.size; i++ {
		if d.array[i] == item {
			index = int(i)
			break
		}
	}

	return index
}

func (d *dynamicArray) remove(item int) {
	found := d.find(item)
	for found != -1 {
		d.delete(uint(found))
		found = d.find(item)
	}
}

func (d *dynamicArray) resize(newCapacity uint) {
	d.capacity = newCapacity
	if d.size > d.capacity {
		d.size = d.capacity
	}

	newArray := make([]int, d.capacity, d.capacity)
	var i uint
	for i = 0; i < d.size; i++ {
		newArray[i] = d.array[i]
	}

	d.array = newArray
}

func main() {}
