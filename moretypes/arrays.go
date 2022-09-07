package main

import "fmt"

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
}
