package main

import "fmt"

func main() {

	var a rune

	a = '\u2665'

	fmt.Printf("%c\n", a)
	fmt.Printf("%T\n", a)

}