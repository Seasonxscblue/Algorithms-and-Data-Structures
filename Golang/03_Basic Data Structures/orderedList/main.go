package main

import (
	"fmt"
	"math/rand"
	"time"

	"cn.edu.xtu/orderedList"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	l := orderedList.OrderedList[int]{}

	for i := 0; i < rand.Intn(20)+1; i++ {
		l.Add(rand.Intn(100))
	}

	fmt.Println(l.String())

}
