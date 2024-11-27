package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64

	name = "Calvin"
	age = 22
	height = 1.70

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)	

	fmt.Printf("My name is %s, I'm %d years old and I'm %.2fm.\n", name, age, height)
}