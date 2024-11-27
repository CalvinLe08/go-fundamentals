package main

import "fmt"


func increment(num int) {
	num = num + 5

	fmt.Println(num)
}

func main() {
	x := 5

	increment(x)
	
	fmt.Println(x)
}