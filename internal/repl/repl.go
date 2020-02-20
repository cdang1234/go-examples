package repl

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/cdang1234/go-examples/internal/transaction"
	"github.com/eiannone/keyboard"
)

type Processor struct {
	KeyMap   map[string]int // maps key to value
	ValueMap map[int]int    // tracks number of keys to a value
}

func (r *Processor) Run() {
	// initialize keyboard stroke detector
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	fmt.Print(">")

	// event log
	queue := list.New()
	str := ""

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			// break if user presses ESC key
			break
		} else if key == keyboard.KeyEnter {
			// start new line when user presses ENTER key
			fmt.Print("\n")
			queue.PushBack(r.processEvent(str))
			str = ""
			fmt.Print(">")
		} else if key == keyboard.KeySpace {
			// append white space
			str += " "
			fmt.Print(" ")
		} else if key == keyboard.KeyBackspace2 {
			// remove last character from str and print modified string on new line
			if len(str) > 0 {
				str = str[:len(str)-1]
				fmt.Printf("\n>%v", str)
			}
		} else {
			// store command
			next := strings.ToUpper(string(char))
			str += next
			fmt.Print(next)
		}
	}
}

func (r *Processor) processEvent(event string) transaction.Event {
	components := strings.Split(event, " ")

	switch components[0] {
	case "SET":
		if !validateInput(components, 3) {
			return transaction.Event{}
		}

		value, err := strconv.Atoi(components[2])
		if err != nil {
			panic(err)
		}
		// case where entry exists in key map --> must decrement old value map
		if _, ok := r.KeyMap[components[1]]; ok {
			oldValue := r.KeyMap[components[1]]
			r.ValueMap[oldValue]--
		}

		// set new key to value mapping
		r.KeyMap[components[1]] = value

		// if value already exists in value map, then increment
		if _, ok := r.ValueMap[value]; ok {
			r.ValueMap[value]++
		} else {
			// new value. must init in value map with value of one
			r.ValueMap[value] = 1
		}
	case "GET":
		if !validateInput(components, 2) {
			return transaction.Event{}
		}

		// get value from key map
		fmt.Println(r.KeyMap[components[1]])
	case "DELETE":
		if !validateInput(components, 2) {
			return transaction.Event{}
		}

		// update value map
		r.ValueMap[r.KeyMap[components[1]]]--

		// update key map
		delete(r.KeyMap, components[1])
	case "COUNT":
		if !validateInput(components, 2) {
			return transaction.Event{}
		}

		value, err := strconv.Atoi(components[1])
		if err != nil {
			panic(err)
		}
		// leverage value map to return number of keys with value
		fmt.Println(r.ValueMap[value])
	}
	return transaction.Event{}
}

func validateInput(components []string, argsCount int) bool {
	if len(components) != argsCount || len(components[1]) == 0 {
		fmt.Println("invalid input")
		return false
	}
	return true
}

// func Commit() {

// }

// func Rollback() {

// }

// func Begin() {

// }
