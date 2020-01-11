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
