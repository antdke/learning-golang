package main

import "fmt"

// taking param variable and assigning new value in function
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	// val doesn't change
	zeroval(i)
	fmt.Println("zeroval:", i)

	// pointer changes because you're assigning it to a different memory address
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// printing a pointer
	fmt.Println("pointer:", &i)
}
