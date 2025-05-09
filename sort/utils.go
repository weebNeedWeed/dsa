package main

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
