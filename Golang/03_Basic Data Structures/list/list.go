package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 结点结构体
type node struct {
	data int
	next *node
}

func newNode(data int, next *node) *node {
	n := new(node)
	n.data = data
	n.next = next
	return n
}

// 链表结构体
type List struct {
	head *node
	size int
}

func (l *List) Clear() {
	l.size = 0
	l.head = nil
}

func (l *List) Get(indx int) int {
	return l.node(indx).data
}

func (l *List) Set(indx int, data int) {
	node := l.node(indx)
	node.data = data
}

func (l *List) Insert(indx, v int) {
	l.rangeCheckForAdd(indx)

	if indx == 0 {
		l.head = newNode(v, l.head)
	} else {
		prev := l.node(indx - 1)
		prev.next = newNode(v, prev.next)
	}
	l.size++
}

func (l *List) Append(data int) {
	l.Insert(l.size, data)
}

func (l *List) Remove(indx int) {
	if indx == 0 {
		l.head = l.head.next
	} else {
		prev := l.node(indx - 1)
		prev.next = prev.next.next
	}

	l.size--
}

func (l *List) Contains(data int) bool {
	return l.IndexOf(data) != -1
}

func (l *List) IndexOf(data int) int {
	node := l.head
	for i := 0; i < l.size; i++ {
		if data == node.data {
			return i
		}
		node = node.next
	}

	return -1
}

func (l *List) Len() int {
	return l.size
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

func (l *List) node(indx int) *node {
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

func (l *List) rangeCheck(indx int) error {
	if indx < 0 || indx >= l.size {
		return errors.New("list subscript out of range")
	}
	return nil
}

func (l *List) rangeCheckForAdd(indx int) error {
	if indx < 0 || indx > l.size {
		return errors.New("list subscript out of range")
	}
	return nil
}

func (l *List) String() string {
	str := ""
	n := l.head
	for i := 0; i < l.size; i++ {
		if i != 0 {
			str += "->"
		}
		str += strconv.Itoa(n.data)
		n = n.next
	}

	return str
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var l List
	l.Append(2)
	l.Append(3)
	l.Append(5)
	l.Append(8)
	l.Append(9)
	l.Append(10)
	l.Append(18)
	l.Append(26)
	l.Append(32)

	// Get
	r := rand.Intn(l.Len())
	fmt.Printf("l.Get(%d) = %d\n", r, l.Get(r))

	// Set
	r1, r2 := rand.Intn(l.Len()), rand.Intn(100)+1
	fmt.Printf("l.Set(%d, %d): \n", r1, r2)
	l.Set(r1, r2)
	fmt.Println(l.String())

	// Insert
	// 头插
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
	r = rand.Intn(l.Len())
	fmt.Printf("l.Remove(%d)\n", r)
	l.Remove(r)
	fmt.Println(l.String())

	// Contains
	r = rand.Intn(rand.Intn(100))
	fmt.Printf("l.Contains(%d) = %v\n", r, l.Contains(r))

	// Indexof
	r = rand.Intn(l.Len())
	fmt.Printf("l.IndexOf(%d) = %d\n", l.Get(r), l.IndexOf(l.Get(r)))

	// Len
	fmt.Printf("l.Len() = %d\n", l.Len())

	// Clear
	l.Clear()
	fmt.Printf("L.Clear()\n")

	// IsEmpty()
	fmt.Printf("l.IsEmpty() = %v\n", l.IsEmpty())
}
