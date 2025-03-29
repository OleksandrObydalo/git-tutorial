package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//fmt.Printf("Тип: %T Значення: %v\n", ToBe, ToBe)
	//fmt.Printf("Тип: %T Значення: %v\n", MaxInt, MaxInt)
	//fmt.Printf("Тип: %T Значення: %v\n", z, z)

	//var i int
	//var f float64
	//var b bool
	//var s string
	//var c complex128
	//fmt.Printf("%v %v %v %q %v\n", i, f, b, s, c)

	//var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	//fmt.Println(x, y, z)

	//var n int = 1
	//var f float64 = float64(n / 2)
	//fmt.Println(n, f)

	//v := 42.5 // зміни мене!
	//fmt.Printf("v має тип %T\n", v)

	//const Pi = math.Pi
	//r := 10.0
	//fmt.Println("Circle with radius r =", r, "has area S =", math.Pi*r*r)

	const (
		// Створіть величезне число, зсунувши на 1 біт вліво на 100 позицій.
		// Іншими словами число 1, за яким слідує 100 нулів.
		Big = 1 << 100
		// Зсуваємо його знову вправо на 99 позицій, так що отримуємо 1<<1, або просто 2.
		Small = Big >> 99
	)

	fmt.Println(Small)
	fmt.Println(Big * 0.1)
	fmt.Println("Version 2")
}
