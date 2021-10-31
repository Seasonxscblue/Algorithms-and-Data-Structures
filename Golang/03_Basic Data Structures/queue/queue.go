package main

import "fmt"

type queue[T any] []T

func (q *queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue[T]) Offer(element T) {
	*q = append(*q, element)
}

func (q *queue[T]) Poll() T {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue[T]) Front() T {
	return (*q)[0]
}

func (q *queue[T]) Size() int {
	return len(*q)
}

func main() {
	fmt.Println(hotPotato([]string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}, 7))
}

func hotPotato(namelist []string, num int) string {
	simQueue := queue[string]{}
	for _, name := range namelist {
		simQueue.Offer(name)
	}

	for simQueue.Size() > 1 {
		for i := 0; i < 7; i++ {
			simQueue.Offer(simQueue.Poll())
		}
		simQueue.Poll()
	}

	return simQueue.Poll()
}
