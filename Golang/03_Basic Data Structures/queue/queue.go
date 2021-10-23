package main

import "fmt"

type Queue struct {
	queue []int
}

func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) offer(element int) {
	q.queue = append(q.queue, element)
}

func (q *Queue) poll() int {
	head := q.queue[0]
	q.queue = q.queue[1:]
	return head
}

func (q *Queue) front() int {
	return q.queue[0]
}

func (q *Queue) size() int {
	return len(q.queue)
}

func main() {
	q := Queue{queue: []int{6, 3, 1, 7, 5, 8, 9, 2, 4}}

	for {
		fmt.Printf("%d", q.poll())
		q.offer(q.poll())

		if q.size() == 1 {
			fmt.Printf("%d", q.poll())
			break
		}
	}
}
