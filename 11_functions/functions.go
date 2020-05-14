package main

import "fmt"

// a function to create a sum of 2 numbers
func getSum(a int, b int) int {
	return a + b
}

// a function to cubed of a number
func plusPlus(a, b, c int) int {

	return a + b + c
}

// RETURNING MULTIPLE RETURNS

// Go allows for returning multiple values
// "int, int" means that we're expecting 2 ints
func vals() (int, int) {
	return 3, 7
}

// VARIADIC FUNCTIONS

// this function takes an arbitrary amount of params
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// CLOSURES

func main() {
	n := getSum(1, 2)
	fmt.Println("1 + 2 =", n)

	m := plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", m)

	a, b := vals()
	fmt.Println("First value:", a)
	fmt.Println("Second value:", b)

	// if you only want a subset of the returns use the blank identier _
	_, c := vals()
	fmt.Println("This is only the 2nd value:", c)

	// printing a variadic function
	sum(1, 2)
	sum(1, 2, 3)
	// slice can hold x number of elements, great to run through a variadic function
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
