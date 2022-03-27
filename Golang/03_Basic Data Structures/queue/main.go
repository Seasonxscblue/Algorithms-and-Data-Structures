package main

import (
	"fmt"

	"cn.edu.xtu/queue"
)

func main() {
	fmt.Println(hotPotato([]string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}, 7))
}

func hotPotato(namelist []string, num int) string {
	simQueue := queue.Queue[string]{}
	for _, name := range namelist {
		simQueue.Offer(name)
	}

	for simQueue.Size() > 1 {
		for i := 0; i < 7; i++ {
			simQueue.Offer(simQueue.Poll())
		}
		simQueue.Poll()
	}

	return simQueue.Poll()
}
