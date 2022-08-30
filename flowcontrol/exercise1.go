/*
As a way to play with functions and loops, let's implement a square root function:
given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop.
Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)

Repeating this adjustment makes the guess better and better until we reach an
answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided.
A decent starting guess for z is 1, no matter what the input.
To begin with, repeat the calculation 10 times and print each z along the way.
See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.
*/

package main

import (
	"fmt"
	"math"
)

const tolerance = 1e-9 // scientific notation!

func Sqrt(x float64) float64 {
	z := 1.0
	for diff := math.Abs(z*z - x); diff > tolerance; diff = math.Abs(z*z - x) {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, x, diff)
	}
	return z
}

func SqrtWithIterCondition(x float64) float64 {
	z := 1.0
	zPrev := 0.0
	for diff := math.Abs(z - zPrev); diff > tolerance; diff = math.Abs(z - zPrev) {
		zPrev = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, x, diff)
	}
	return z
}

func main() {
	fmt.Println("Found result:", Sqrt(2))
	fmt.Println("Found result:", SqrtWithIterCondition(2))
}
