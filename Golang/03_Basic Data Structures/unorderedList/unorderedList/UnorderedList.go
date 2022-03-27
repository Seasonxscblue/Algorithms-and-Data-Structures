package unorderedList

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
type UnorderedList[T constraints.Ordered] struct {
	head *node[T]
	size int
}

func (l *UnorderedList[T]) Clear() {
	l.size = 0
	l.head = nil
}

func (l *UnorderedList[T]) Get(indx int) T {
	return l.node(indx).data
}

func (l *UnorderedList[T]) Set(indx int, data T) {
	node := l.node(indx)
	node.data = data
}

func (l *UnorderedList[T]) Insert(indx int, v T) {
	l.rangeCheckForAdd(indx)

	if indx == 0 {
		l.head = newNode(v, l.head)
	} else {
		prev := l.node(indx - 1)
		prev.next = newNode(v, prev.next)
	}
	l.size++
}

func (l *UnorderedList[T]) Append(data T) {
	l.Insert(l.size, data)
}

func (l *UnorderedList[T]) Remove(indx int) {
	if indx == 0 {
		l.head = l.head.next
	} else {
		prev := l.node(indx - 1)
		prev.next = prev.next.next
	}

	l.size--
}

func (l *UnorderedList[T]) Contains(data T) bool {
	return l.IndexOf(data) != -1
}

func (l *UnorderedList[T]) IndexOf(data T) int {
	node := l.head
	for i := 0; i < l.size; i++ {
		if data == node.data {
			return i
		}
		node = node.next
	}

	return -1
}

func (l *UnorderedList[T]) Len() int {
	return l.size
}

func (l *UnorderedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *UnorderedList[T]) node(indx int) *node[T] {
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

func (l *UnorderedList[T]) rangeCheck(indx int) error {
	if indx < 0 || indx >= l.size {
		return errors.New("UnorderedList subscript out of range")
	}
	return nil
}

func (l *UnorderedList[T]) rangeCheckForAdd(indx int) error {
	if indx < 0 || indx > l.size {
		return errors.New("UnorderedList subscript out of range")
	}
	return nil
}

func (l *UnorderedList[T]) String() string {
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
