package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

const ConvertString = "0123456789ABCDEF"

func listsum[T constraints.Integer | constraints.Float](numList []T) T {
	if len(numList) == 1 {
		return numList[0]
	}

	return numList[0] + listsum(numList[1:])
}

func toStr(n, base int) string {
	if n < base {
		return string(ConvertString[n])
	}

	return toStr(n/base, base) + string(ConvertString[n%base])
}

func main() {
	fmt.Println(toStr(11, 16))
}
