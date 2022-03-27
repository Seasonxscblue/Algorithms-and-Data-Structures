package deque

type Deque[T any] []T

func (d *Deque[T]) IsEmpty() bool {
	return len(*d) == 0
}

func (d *Deque[T]) AddRear(item T) {
	*d = append([]T{item}, *d...)
}

func (d *Deque[T]) AddFront(item T) {
	*d = append(*d, item)
}

func (d *Deque[T]) RemoveFront() T {
	head := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return head
}

func (d *Deque[T]) RemoveRear() T {
	tail := (*d)[0]
	*d = (*d)[1:]
	return tail
}

func (d *Deque[T]) Size() int {
	return len(*d)
}
