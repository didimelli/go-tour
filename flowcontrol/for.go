package main

import "fmt"

func sum(sum int) int {
	for i := 0; i < 10; i++ {
		sum += 1
	}
	return sum
}

// init and post are optional!
func sumOptional(sum int) int {
	for sum < 1000 {
		sum += sum
	}
	return sum
}

// used as a while
func sumWhileLike(sum int) int {
	for sum < 1000 {
		sum += sum
	}
	return sum
}

// infinite
func infinite() {
	for {
	}
}

func main() {
	fmt.Println(sum(10))
	fmt.Println(sumOptional(1))
	fmt.Println(sumWhileLike(1))
	// loops indefinitely
	fmt.Println("Warning, this will loop indifinitely.")
	infinite()
}
