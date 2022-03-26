package main

import (
	//"errors"
	"fmt"
	//"os"
	//"math/rand"
	//"time"
)

type calculable interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string
}

//// 结点结构体
//type node[T calculable] struct {
//	data T
//	next *node[T]
//}
//
//func newNode[T calculable](data T, next *node[T]) *node[T] {
//	n := new(node[T])
//	n.data, n.next = data, next
//	return n
//}

func compare[T calculable] (o1, o2 T) T {
	return o1 - o2
}

//// 链表结构体
//type orderedList[T comparable] struct {
//	head *node[T]
//	size int
//}
//
//func (l *orderedList[T]) Clear() {
//	l.size = 0
//	l.head = nil
//}
//
//func (l *orderedList[T]) Get(indx int) T {
//	l.rangeCheck(indx)
//	return l.node(indx).data
//}
//
////func (l *orderedList[T]) Insert(indx int, v T) {
////	l.rangeCheckForOperation(indx)
////
////	if indx == 0 {
////		l.head = newNode(v, l.head)
////	} else {
////		prev := l.node(indx - 1)
////		prev.next = newNode(v, prev.next)
////	}
////	l.size++
////}
////
//
//func (l *orderedList[T]) Add(data T) {
//
//}
//
//func (l *orderedList[T]) Remove(indx int) {
//	l.rangeCheck(indx)
//
//	if indx == 0 {
//		l.head = l.head.next
//	} else {
//		prev := l.node(indx - 1)
//		prev.next = prev.next.next
//	}
//
//	l.size--
//}
//
//func (l *orderedList[T]) Contains(data T) bool {
//	return l.IndexOf(data) != -1
//}
//
//func (l *orderedList[T]) IndexOf(data T) int {
//	node := l.head
//	stop := false
//
//	for i := 0; i < l.size && !stop; i++ {
//		if data == node.data {
//			return i
//		} else {
//			if node.data > data {
//				stop = true
//			} else {
//				node = node.next
//			}
//		}
//	}
//	return -1
//}
//
//func (l *orderedList[T]) Len() int {
//	return l.size
//}
//
//func (l *orderedList[T]) IsEmpty() bool {
//	return l.size == 0
//}
//
//func (l *orderedList[T]) node(indx int) *node[T] {
//	err := l.rangeCheck(indx)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	n := l.head
//	for i := 0; i < indx; i++ {
//		n = n.next
//	}
//	return n
//}
//
//func (l *orderedList[T]) rangeCheck(indx int) error {
//	if indx < 0 || indx >= l.size {
//		return errors.New("orderedList subscript out of range")
//	}
//	return nil
//}
//
//func (l *orderedList[T]) rangeCheckForOperation(indx int) error {
//	if indx < 0 || indx > l.size {
//		return errors.New("orderedList subscript out of range")
//	}
//	return nil
//}
//
//func (l *orderedList[T]) String() string {
//	str := ""
//	n := l.head
//	for i := 0; i < l.size; i++ {
//		if i != 0 {
//			str += "->"
//		}
//		str += fmt.Sprint(n.data)
//		n = n.next
//	}
//
//	return str
//}

func main() {
	//rand.Seed(time.Now().UnixNano())
	//
	//l := orderedList[int]{}
	//l.Append(2)
	//l.Append(3)
	//l.Append(5)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	//l.Append(18)
	//l.Append(26)
	//l.Append(32)
	//
	//fmt.Println(l.String())
	//// Get
	//fmt.Println("==========Get==========")
	//r := rand.Intn(l.Len())
	//fmt.Printf("l.Get(%d) = %d\n", r, l.Get(r))
	//
	//// Set
	//fmt.Println("==========Set==========")
	//r1, r2 := rand.Intn(l.Len()), rand.Intn(100)+1
	//fmt.Printf("l.Set(%d, %d): \n", r1, r2)
	//l.Set(r1, r2)
	//fmt.Println(l.String())
	//
	//// Insert
	//// 头插
	//fmt.Println("==========Insert==========")
	//r = rand.Intn(100) + 1
	//fmt.Printf("l.Insert(0, %d)\n", r)
	//l.Insert(0, r)
	//fmt.Println(l.String())
	//
	//// 尾插
	//r = rand.Intn(100) + 1
	//fmt.Printf("l.Insert(l.Len, %d)\n", r)
	//l.Insert(l.Len(), r)
	//fmt.Println(l.String())
	//
	//// 中间插
	//r1, r2 = rand.Intn(l.Len())+1, rand.Intn(100)+1
	//fmt.Printf("l.Insert(%d, %d)\n", r1, r2)
	//l.Insert(r1, r2)
	//fmt.Println(l.String())
	//
	//// Remove
	//fmt.Println("==========Remove==========")
	//r = rand.Intn(l.Len())
	//fmt.Printf("l.Remove(%d)\n", r)
	//l.Remove(r)
	//fmt.Println(l.String())
	//
	//// Contains
	//fmt.Println("==========Contains==========")
	//r = rand.Intn(rand.Intn(100))
	//fmt.Printf("l.Contains(%d) = %v\n", r, l.Contains(r))
	//
	//// Indexof
	//fmt.Println("==========Indexof==========")
	//r = rand.Intn(l.Len())
	//fmt.Printf("l.IndexOf(%d) = %d\n", l.Get(r), l.IndexOf(l.Get(r)))
	//
	//// Len
	//fmt.Println("==========Len==========")
	//fmt.Printf("l.Len() = %d\n", l.Len())
	//
	//// Clear
	//fmt.Println("==========Clear==========")
	//l.Clear()
	//fmt.Printf("L.Clear()\n")
	//
	//// IsEmpty()
	//fmt.Println("==========IsEmpty==========")
	//fmt.Printf("l.IsEmpty() = %v\n", l.IsEmpty())

	fmt.Println(compare(3, 4))
}
