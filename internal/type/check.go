package check

import (
	"fmt"
	"reflect"
)

type Custom struct {
	value int
}

func CustomFindType() {
	c := Custom{value: 0}

	ct := reflect.TypeOf(c).String()
	fmt.Println(ct == "main.Custom") // returns true
	fmt.Printf("%T: %s\n", ct, ct)
}

func SimpleFindType() {
	x := 42
	y := float32(43.3)
	z := "hello"

	xt := reflect.TypeOf(x).Kind()
	yt := reflect.TypeOf(y).Kind()
	zt := reflect.TypeOf(z).Kind()

	fmt.Printf("%T: %s\n", xt, xt)
	fmt.Printf("%T: %s\n", yt, yt)
	fmt.Printf("%T: %s\n", zt, zt)

	if xt == reflect.Int {
		println(">> x is int")
	}
	if yt == reflect.Float32 {
		println(">> y is float32")
	}
	if zt == reflect.String {
		println(">> z is string")
	}
}
