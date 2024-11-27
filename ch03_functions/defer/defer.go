package main

import "fmt"

func main() {

	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
	}
}