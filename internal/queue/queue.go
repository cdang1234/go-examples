package queue

import (
	"container/list"
	"fmt"
)

// https://golang.org/pkg/container/list/#List.PushBackList

func Queue() {
	queue := list.New()
	queue.PushBack("hello")
	queue.PushBack(" world")

	for queue.Len() > 0 {
		e := queue.Front() // First element
		fmt.Print(e.Value)

		queue.Remove(e) // Dequeue
	}
}
