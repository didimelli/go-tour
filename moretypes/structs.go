package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Printf("%T", Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4 // dot notation!
	// you can also access through pointer (but directly!!)
	p := &v
	p.X = 1e9
	fmt.Println(v)
	v1 := Vertex{}     // defaults
	v2 := Vertex{X: 1} // named
	fmt.Println(v1, v2)
}
