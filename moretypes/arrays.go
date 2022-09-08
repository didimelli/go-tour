package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	fmt.Println(a)
	a[0] = "asd"
	a[1] = "ciao"
	fmt.Println(a, a[0], a[1])
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are references to underlying array!
	names := [4]string{
		"a",
		"b",
		"c",
		"d",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(a, b)

	c[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literal with structs!
	d := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(d)

	// slices works as in python, you can omit high/low bound
	e := []int{2, 3, 5, 7, 11, 13}

	e = e[1:4]
	fmt.Println(e)

	e = e[:2]
	fmt.Println(e)

	e = e[1:]
	fmt.Println(e)

	// slices has lenght and capacity!
	fmt.Println("length:", len(e), "capacity:", cap(e))

	// nil slices!
	var n []int
	fmt.Println(n, len(n), cap(n))
	if n == nil {
		fmt.Println("nil!")
	}

	// make to create dynamically sized arrays!
	f := make([]int, 5)    // len is 5
	g := make([]int, 0, 5) // len 0, capacity 5
	fmt.Println(f, g)

	// slice of slices
	// tic tac toe
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// players
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var og []int
	printSlice(og)
	// append works on nil
	og = append(og, 0)
	printSlice(og)
	// slice grows as needed
	og = append(og, 1)
	printSlice(og)
	// append more at a time
	og = append(og, 2, 3, 4, 5)
	printSlice(og)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
