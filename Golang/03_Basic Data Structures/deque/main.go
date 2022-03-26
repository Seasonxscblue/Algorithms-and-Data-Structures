package main

import (
	"fmt"

	"cn.edu.xtu/Deque"
)

func main() {
	fmt.Println(palchecker("madam"))
}

func palchecker(aString string) bool {
	charDeque := Deque.Deque[rune]{}

	for _, ch := range aString {
		charDeque.AddRear(ch)
	}

	stillEqual := true
	for charDeque.Size() > 1 && stillEqual {
		if first, last := charDeque.RemoveFront(), charDeque.RemoveRear(); first != last {
			stillEqual = false
		}
	}
	return stillEqual
}
