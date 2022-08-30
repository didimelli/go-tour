package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// numeric constants are high precision values
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() { // looks like go does not like two main func in the same folder (aka module?) -> "main redeclared in this block"
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) // with %T i print the type of the var
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i int     // zero values assigned by default if no initial value is provided
	var f float64 // zero values assigned by default if no initial value is provided
	var b bool    // zero values assigned by default if no initial value is provided
	var s string  // zero values assigned by default if no initial value is provided
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y)) // type casting!
	var z uint = uint(ff)
	fmt.Println(x, y, z, ff)
	// type inference
	// j := i // j is int since i is int
	// const works as var, but cannot be declared with :=
	x = 4
	const c = 2
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) // this breakes because it overflows
}
