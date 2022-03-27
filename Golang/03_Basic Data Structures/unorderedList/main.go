package main

import (
	"fmt"
	"math/rand"
	"time"

	"cn.edu.xtu/unorderedList"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	l := unorderedList.UnorderedList[int]{}
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
	fmt.Printf("\n==========Get==========\n")
	r := rand.Intn(l.Len())
	fmt.Printf("l.Get(%d) = %d\n", r, l.Get(r))

	// Set
	fmt.Printf("\n==========Set==========\n")
	r1, r2 := rand.Intn(l.Len()), rand.Intn(100)+1
	fmt.Printf("l.Set(%d, %d): \n", r1, r2)
	l.Set(r1, r2)
	fmt.Println(l.String())

	// Insert
	// 头插
	fmt.Printf("\n==========Insert==========\n")
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
	fmt.Printf("\n==========Remove==========\n")
	r = rand.Intn(l.Len())
	fmt.Printf("l.Remove(%d)\n", r)
	l.Remove(r)
	fmt.Println(l.String())

	// Contains
	fmt.Printf("\n==========Contains==========\n")
	r = rand.Intn(rand.Intn(100))
	fmt.Printf("l.Contains(%d) = %v\n", r, l.Contains(r))

	// Indexof
	fmt.Printf("\n==========Indexof==========\n")
	r = rand.Intn(l.Len())
	fmt.Printf("l.IndexOf(%d) = %d\n", l.Get(r), l.IndexOf(l.Get(r)))

	// Len
	fmt.Printf("\n==========Len==========\n")
	fmt.Printf("l.Len() = %d\n", l.Len())

	// Clear
	fmt.Printf("\n==========Clear==========\n")
	l.Clear()
	fmt.Printf("l.Clear()\n")

	// IsEmpty()
	fmt.Printf("\n==========IsEmpty==========\n")
	fmt.Printf("l.IsEmpty() = %v\n", l.IsEmpty())
}
