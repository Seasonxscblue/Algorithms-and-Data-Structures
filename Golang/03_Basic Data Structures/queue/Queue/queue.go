package Queue

type Queue[T any] []T

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Offer(element T) {
	*q = append(*q, element)
}

func (q *Queue[T]) Poll() T {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue[T]) Front() T {
	return (*q)[0]
}

func (q *Queue[T]) Size() int {
	return len(*q)
}
