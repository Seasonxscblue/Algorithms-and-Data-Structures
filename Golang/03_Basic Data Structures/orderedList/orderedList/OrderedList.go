package orderedList

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
)

// 结点结构体
type node[T constraints.Ordered] struct {
	data T
	next *node[T]
}

func newNode[T constraints.Ordered](data T, next *node[T]) *node[T] {
	n := new(node[T])
	n.data, n.next = data, next
	return n
}

// 链表结构体
type OrderedList[T constraints.Ordered] struct {
	head *node[T]
	size int
}

func (l *OrderedList[T]) Clear() {
	l.size = 0
	l.head = nil
}

func (l *OrderedList[T]) Get(indx int) T {
	l.rangeCheck(indx)
	return l.node(indx).data
}

func (l *OrderedList[T]) Add(data T) {
	var i int

	if l.size == 0 {
		l.head = newNode(data, nil)
	} else {
		for i = 0; i < l.size; i++ {
			if data <= l.node(i).data {
				if i == 0 {
					l.head = newNode(data, l.head)
				} else {
					l.node(i - 1).next = newNode(data, l.node(i))
				}
				break
			}
		}
		if i == l.size {
			l.node(l.size - 1).next = newNode(data, nil)
		}
	}
	l.size++
}

func (l *OrderedList[T]) Remove(indx int) {
	if indx == 0 {
		l.head = l.head.next
	} else {
		prev := l.node(indx - 1)
		prev.next = prev.next.next
	}

	l.size--
}

func (l *OrderedList[T]) Contains(data T) bool {
	return l.IndexOf(data) != -1
}

func (l *OrderedList[T]) IndexOf(data T) int {
	lo := 0
	hi := l.size - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		n := l.node(mid)
		if n.data == data {
			return mid
		} else if n.data > data {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

func (l *OrderedList[T]) Len() int {
	return l.size
}

func (l *OrderedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *OrderedList[T]) node(indx int) *node[T] {
	err := l.rangeCheck(indx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n := l.head
	for i := 0; i < indx; i++ {
		n = n.next
	}
	return n
}

func (l *OrderedList[T]) rangeCheck(indx int) error {
	if indx < 0 || indx >= l.size {
		return errors.New("orderedList subscript out of range")
	}
	return nil
}

func (l *OrderedList[T]) String() string {
	str := ""

	n := l.head
	for i := 0; i < l.size; i++ {
		if i != 0 {
			str += " -> "
		}
		str += fmt.Sprint(n.data)
		n = n.next
	}

	return str
}
