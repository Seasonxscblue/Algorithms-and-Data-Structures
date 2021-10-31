package main

import (
	"errors"
	"fmt"
	"os"
	"math/rand"
	"time"
)

// 结点结构体
type node[T comparable] struct {
	data T
	next *node[T]
}

func newNode[T comparable](data T, next *node[T]) *node[T] {
	n := new(node[T])
	n.data, n.next = data, next
	return n
}

// 链表结构体
type list[T comparable] struct {
	head *node[T]
	size int
}

func (l *list[T]) Clear() {
	l.size = 0
	l.head = nil
}

func (l *list[T]) Get(indx int) T {
	return l.node(indx).data
}

func (l *list[T]) Set(indx int, data T) {
	node := l.node(indx)
	node.data = data
}

func (l *list[T]) Insert(indx int, v T) {
	l.rangeCheckForAdd(indx)

	if indx == 0 {
		l.head = newNode(v, l.head)
	} else {
		prev := l.node(indx - 1)
		prev.next = newNode(v, prev.next)
	}
	l.size++
}

func (l *list[T]) Append(data T) {
	l.Insert(l.size, data)
}

func (l *list[T]) Remove(indx int) {
	if indx == 0 {
		l.head = l.head.next
	} else {
		prev := l.node(indx - 1)
		prev.next = prev.next.next
	}

	l.size--
}

func (l *list[T]) Contains(data T) bool {
	return l.IndexOf(data) != -1
}

func (l *list[T]) IndexOf(data T) int {
	node := l.head
	for i := 0; i < l.size; i++ {
		if data == node.data {
			return i
		}
		node = node.next
	}

	return -1
}

func (l *list[T]) Len() int {
	return l.size
}

func (l *list[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *list[T]) node(indx int) *node[T] {
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

func (l *list[T]) rangeCheck(indx int) error {
	if indx < 0 || indx >= l.size {
		return errors.New("list subscript out of range")
	}
	return nil
}

func (l *list[T]) rangeCheckForAdd(indx int) error {
	if indx < 0 || indx > l.size {
		return errors.New("list subscript out of range")
	}
	return nil
}

func (l *list[T]) String() string {
	str := ""
	n := l.head
	for i := 0; i < l.size; i++ {
		if i != 0 {
			str += "->"
		}
		str += fmt.Sprint(n.data)
		n = n.next
	}

	return str
}

func main() {
	rand.Seed(time.Now().UnixNano())

	l := list[int]{}
	l.Append(2)
	l.Append(3)
	l.Append(5)
	l.Append(8)
	l.Append(9)
	l.Append(10)
	l.Append(18)
	l.Append(26)
	l.Append(32)

	fmt.Println(l.String())
	// Get
	fmt.Println("==========Get==========")
	r := rand.Intn(l.Len())
	fmt.Printf("l.Get(%d) = %d\n", r, l.Get(r))

	// Set
	fmt.Println("==========Set==========")
	r1, r2 := rand.Intn(l.Len()), rand.Intn(100)+1
	fmt.Printf("l.Set(%d, %d): \n", r1, r2)
	l.Set(r1, r2)
	fmt.Println(l.String())

	// Insert
	// 头插
	fmt.Println("==========Insert==========")
	r = rand.Intn(100) + 1
	fmt.Printf("l.Insert(0, %d)\n", r)
	l.Insert(0, r)
	fmt.Println(l.String())

	// 尾插
	r = rand.Intn(100) + 1
	fmt.Printf("l.Insert(l.Len, %d)\n", r)
	l.Insert(l.Len(), r)
	fmt.Println(l.String())

	// 中间插
	r1, r2 = rand.Intn(l.Len())+1, rand.Intn(100)+1
	fmt.Printf("l.Insert(%d, %d)\n", r1, r2)
	l.Insert(r1, r2)
	fmt.Println(l.String())

	// Remove
	fmt.Println("==========Remove==========")
	r = rand.Intn(l.Len())
	fmt.Printf("l.Remove(%d)\n", r)
	l.Remove(r)
	fmt.Println(l.String())

	// Contains
	fmt.Println("==========Contains==========")
	r = rand.Intn(rand.Intn(100))
	fmt.Printf("l.Contains(%d) = %v\n", r, l.Contains(r))

	// Indexof
	fmt.Println("==========Indexof==========")
	r = rand.Intn(l.Len())
	fmt.Printf("l.IndexOf(%d) = %d\n", l.Get(r), l.IndexOf(l.Get(r)))

	// Len
	fmt.Println("==========Len==========")
	fmt.Printf("l.Len() = %d\n", l.Len())

	// Clear
	fmt.Println("==========Clear==========")
	l.Clear()
	fmt.Printf("L.Clear()\n")

	// IsEmpty()
	fmt.Println("==========IsEmpty==========")
	fmt.Printf("l.IsEmpty() = %v\n", l.IsEmpty())
}
