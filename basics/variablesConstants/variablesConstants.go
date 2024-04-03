package main

import "fmt"

func main() {

	const pi float64 = 3.14159265359

	fmt.Println(pi)

	// multiple variable declaration

	var (
		varA = 2
		varB = 2
	)

	fmt.Println(varA, varB)

	// Strings are between " or `

	var name string = "Josh Ware"

	fmt.Println(len(name))

}
