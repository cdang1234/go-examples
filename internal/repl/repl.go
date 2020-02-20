package repl

import (
	"container/list"
	"fmt"
	"github.com/eiannone/keyboard"
	"strconv"
	"strings"
)

type Processor struct {
	KeyMap   map[string]int // maps key to value
	ValueMap map[int]int    // tracks number of keys to a value
	Stack    *list.List
}

type Event struct {
	command string
	args    []string
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
			event := r.processEvent(str)
			if event.command == "SET" || event.command == "DELETE" || event.command == "BEGIN" {
				r.Stack.PushFront(event)
			}
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

func (r *Processor) processEvent(event string) Event {
	components := strings.Split(event, " ")

	switch components[0] {
	case "SET":
		if !validateInput(components, 3) {
			return Event{}
		}

		r.Set(components)

		return Event{command: "DELETE", args: []string{components[1]}}
	case "GET":
		if !validateInput(components, 2) {
			return Event{}
		}

		// get value from key map
		fmt.Println(r.KeyMap[components[1]])

		return Event{command: "GET", args: components[1:]}
	case "DELETE":
		if !validateInput(components, 2) {
			return Event{}
		}

		originalKey, originalVal := r.Delete(components)

		return Event{command: "SET", args: []string{originalKey, originalVal}}
	case "COUNT":
		if !validateInput(components, 2) {
			return Event{}
		}

		value, err := strconv.Atoi(components[1])
		if err != nil {
			panic(err)
		}
		// leverage value map to return number of keys with value
		fmt.Println(r.ValueMap[value])

		return Event{command: "COUNT", args: components[1:]}
	case "BEGIN":
		return Event{command: "BEGIN"}
	case "COMMIT":
		for r.Stack.Len() > 0 {
			e := r.Stack.Front()
			r.Stack.Remove(e)
			if e.Value.(Event).command == "BEGIN" {
				break
			}
		}
		return Event{}
	case "ROLLBACK":
		for r.Stack.Len() > 0 {
			e := r.Stack.Front()
			r.Stack.Remove(e)
			if e.Value.(Event).command == "BEGIN" {
				break
			} else {
				// undo event
				if e.Value.(Event).command == "DELETE" {
					r.Delete(e.Value.(Event).args)
				} else {
					r.Set(e.Value.(Event).args)
				}
			}
		}

		return Event{}
	}
	return Event{}
}

func validateInput(components []string, argsCount int) bool {
	if len(components) != argsCount || len(components[1]) == 0 {
		fmt.Println("invalid input")
		return false
	}
	return true
}

func (r *Processor) Delete(components []string) (string, string) {
	value := r.KeyMap[components[1]]
	key := components[1]

	// update value map
	r.ValueMap[value]--

	// update key map
	delete(r.KeyMap, key)

	return key, strconv.Itoa(value)
}

func (r *Processor) Set(components []string) {
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
}
