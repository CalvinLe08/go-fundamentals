package main

import "fmt"

func calculator(a, b int) (mul int, div int, add int, sub int) {
	mul = a * b
	div = a / b
	add = a + b
	sub = a - b

	return mul, div, add, sub
}

func main() {
	a, b, c, d := calculator(1,2)

	fmt.Printf("%v %v %v %v\n", a,b,c,d)
}