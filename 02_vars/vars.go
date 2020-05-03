package main

import "fmt"

func main() {
	// data types are tracked like in TypeScript
	// every variable has to be used
	const name = "Anthony Diké"

	// Shorthand
	firstName, email := "Anthony", "anthonydike@nyu.edu"

	fmt.Println(name)
	fmt.Println("My email is", email)
	fmt.Printf("%T", firstName)
}
