package main

import "fmt"

func main() {

	// Iterating over a slice to get the sum of all values in that slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println("In loop sum:", sum)
	}
	fmt.Println("Out of loop sum:", sum)

	// identifying the index of a certain value in a range
	for i, num := range nums {
		if num == 3 {
			fmt.Println("I found it's index:", i)
		}
	}

	// using range to iterate over a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterating over the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// interating over strings returns its Unicode code points
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
