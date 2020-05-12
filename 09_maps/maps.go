package main

import "fmt"

func main() {

	// declaring a map
	m := make(map[string]int)

	// declaring keys and initializing their values
	m["k1"] = 7
	m["k2"] = 3

	// printing out the map
	fmt.Println("map:", m)

	// grab a value from a key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// getting the length of the map
	fmt.Println("len:", len(m))

	// deleting a key
	delete(m, "k2")
	fmt.Println("new map:", m)

	// checking if a value is present
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize new map on the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
