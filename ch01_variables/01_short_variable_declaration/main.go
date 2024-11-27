package main

import "fmt"

func main() {
	// using type inference

	name := "Khang"
	age := 22
	height := 1.70

	fmt.Printf("My name is %v, I'm %v years old and I'm definitely taller than %.2f cm.\n", name, age, height)

	fmt.Printf("name type: %T\nnage type: %T\nheight type: %T\n", name, age, height)
}