package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d", "e", "f")
	fmt.Println("new:", s)

	// copying slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slicing slices
	l := s[0:2] // a, b
	fmt.Println("sl1:", l)

	l2 := s[2:5] // c. d. e
	fmt.Println("sl12:", l2)

	l3 := s[:4] // a, b, c, d
	fmt.Println("sl3:", l3)

	l4 := s[4:] // e, f
	fmt.Println("sl4:", l4)

	// declare and initialize a slice in a single line
	t := []string{"x", "y", "z"}
	fmt.Println("dcl:", t)

	// two dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
