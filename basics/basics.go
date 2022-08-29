package main

import (
	"fmt"
	"math"
	"math/rand"
)

// function
func add(x int, y int) int {
	return x + y
}

// function but shortened. casing is camelCase
// if function params have same type, you can only write last one
func addShortened(x, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

// named return values!
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// package level var statement (with init)
var i, j int = 1, 2

func main() {
	fmt.Println("Number is", rand.Intn(100))
	fmt.Printf("String formatting %g.\n", math.Sqrt(7)) // string formatting requires Printf
	fmt.Printf("Pi is %g.\n", math.Pi)                  // exported stuff is Capitalized
	fmt.Println("My add function!", add(1, 3))
	fmt.Println("My add function (short version)!", addShortened(1, 3))
	a, b := swap("first", "second")
	fmt.Println(a, b)
	fmt.Println(split(17))
	// function level var statement (with init the type is taken from init value)
	var c, python, java = true, false, "no!"
	k := 3 // same as var k = 3. valid only a function! outside a keyword is needed
	fmt.Println(i, j, k, c, python, java)
}
