package main

import (
	"fmt"
)

// fibonacci is a function that 
// 			returns a function [func(int)]
// 				    that returns an int 
func fibonacci() func(int) int {
	x, y := 0, 1
	return func(int) int {
		x, y = y, x + y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}