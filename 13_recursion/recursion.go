package main

import "fmt"

// function that calls itself until the base condition is met

// factorial example

func fact(n int) int {
	if n == 0 {
		// base case
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
