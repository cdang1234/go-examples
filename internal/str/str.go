package str

import (
	"fmt"
	"strings"
)

func ExtractCharFromString() {
	str := "Hello World"
	fmt.Println(string(str[0]))
}

func GenerateFromStringBuilder() {
	var str strings.Builder

	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}

	fmt.Println(str.String())
}

// range of alphabet in byte form
// if s[i] >= 97 && s[i] <= 122 {
// 	currChar = s[i]
// }

// regexp example
// var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
