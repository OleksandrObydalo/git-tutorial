package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, eps float64) float64 {
	sq := 1.0
	for math.Abs(sq*sq-x) > eps {
		sq -= (sq*sq - x) / (2 * sq)
	}
	return sq
}

func main() {
	fmt.Println(Sqrt(2, 0.0001))
}
