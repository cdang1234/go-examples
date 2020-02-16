package math

import (
	"fmt"
	"math"
	"math/rand"
)

const MaxUint = math.MaxUint32
const MinUint = 0
const MaxInt = math.MaxInt32
const MinInt = math.MinInt32

func GenRandNumWithRange() {
	min := 10
	max := 30
	fmt.Println(rand.Intn(max-min) + min)
}
