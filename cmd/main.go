package main

import "fmt"

import "github.com/cdang1234/go-examples/internal/convert"

import "github.com/cdang1234/go-examples/internal/str"

func main() {
	fmt.Println("Hello World!")
	// convert byte to int
	convert.ByteToInt()

	// convert int to byte
	convert.IntToByte()

	// convert string to byte
	convert.StringToByte()

	// convert byte to string
	convert.ByteToString()

	// convert string to int
	convert.StringToInt()

	// convert int to string
	convert.IntToString()

	// extract character from string
	str.ExtractCharFromString()

	// string builder
	str.GenerateFromStringBuilder()

	// 2D array access

	// reference by value vs reference by pointer

	// DFS implementation

	// BFS implementation

	// Max Heap implementation

	// Priority Queue implementation

	// Stack implementation
}
