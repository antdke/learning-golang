package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now().Weekday()
	fmt.Println(t)

	switch t {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
