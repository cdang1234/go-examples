package stack

import (
	"container/list"
	"fmt"
)

// https://golang.org/pkg/container/list/#List.PushFront

func Stack() {
	stack := list.New()
	stack.PushFront(" world")
	stack.PushFront("hello")

	for stack.Len() > 0 {
		e := stack.Front() // First element
		fmt.Print(e.Value)

		stack.Remove(e) // Dequeue
	}
}
