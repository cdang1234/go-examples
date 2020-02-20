package main

import (
	"container/list"

	"github.com/cdang1234/go-examples/internal/repl"
)

func main() {
	repl := repl.Processor{KeyMap: make(map[string]int), ValueMap: make(map[int]int), Queue: list.New()}
	repl.Run()
}
