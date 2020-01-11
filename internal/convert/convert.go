package convert

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func IntToByte() {
	var num int64 = -123456789

	// convert int64 to []byte
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, num)
	b := buf[:n]
	fmt.Println(b)
}

func ByteToInt() {
	var num int64 = -123456789

	// convert int64 to []byte
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, num)
	b := buf[:n]

	// convert []byte to int64
	x, n := binary.Varint(b)
	fmt.Printf("x is: %v, n is: %v\n", x, n)
}

func ByteToString() {
	var str = "Hello World"
	var mySlice = []byte(str)
	fmt.Println(string(mySlice))
}

func StringToByte() {
	var str = "Hello World"
	fmt.Println([]byte(str)) //	fmt.Println([]rune(str)) will produce the same value
}

func StringToInt() {
	s := "10"
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
}

func IntToString() {
	t := strconv.Itoa(123)
	fmt.Println(t)
}
