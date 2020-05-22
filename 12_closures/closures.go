package main

import "fmt"

// returns the next int in a sequence
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// intSeq returns a function. Passing it to nextInt allows that function to return the function's function's return value (I think?)
	nextInt := intSeq()

	// it updates the return value and it's scoped that function (nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// proof that the value is scoped to the higher order function (i think that's what it's called)
	newInts := intSeq()
	fmt.Println(newInts())
}
