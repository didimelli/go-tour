package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         // points to i
	fmt.Println(*p) // prints what is pointed by p, i
	*p = 21         // modifies i through pointer
	fmt.Println(i)

	p = &j
	*p = *p / 37 // operation through pointer
	fmt.Println(j)
}
