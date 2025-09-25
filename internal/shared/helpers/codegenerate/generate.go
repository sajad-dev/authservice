package codegenerate

import (
	"math"
	"math/rand"
)

var CodeLen float64 = 8

func Generate() int {
	return rand.Intn(9*int(math.Pow(10, CodeLen))) + int(math.Pow(10, CodeLen))
}
