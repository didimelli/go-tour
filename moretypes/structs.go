package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Printf("%T", Vertex{1, 2})
}
