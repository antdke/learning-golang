package main

import "fmt"

// for loops are the only loop structure in Go
func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// classic initial/condition/action for loop
	for j := 7; j <= 10; j++ {
		fmt.Println(j)
	}

	// for loop will go on forever without a break or return at the end of a function
	for {
		fmt.Println("loop")
		break
	}

	// continue goes to the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
