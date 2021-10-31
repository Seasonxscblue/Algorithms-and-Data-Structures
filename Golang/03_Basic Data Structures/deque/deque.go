package main

import "fmt"

type deque[T any] []T

func (d *deque[T]) IsEmpty() bool {
	return len(*d) == 0
}

func (d *deque[T]) AddRear(item T) {
	*d = append([]T{item}, *d...)
}

func (d *deque[T]) AddFront(item T) {
	*d = append(*d, item)
}

func (d *deque[T]) RemoveFront() T {
	head := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return head
}

func (d *deque[T]) RemoveRear() T {
	tail := (*d)[0]
	*d = (*d)[1:]
	return tail
}

func (d *deque[T]) Size() int {
	return len(*d)
}

func main() {
	fmt.Println(palchecker("madam"))
}

func palchecker(aString string) bool {
	charDeque := deque[rune]{}

	for _, ch := range aString {
		charDeque.AddRear(ch)
	}

	stillEqual := true
	for charDeque.Size() > 1 && stillEqual {
		if first, last := charDeque.RemoveFront(), charDeque.RemoveRear(); first != last {
			stillEqual = false
		}
	}
	return stillEqual
}
