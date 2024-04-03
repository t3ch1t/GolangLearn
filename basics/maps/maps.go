package main

import "fmt"

func main() {

	student := make(map[string]int)

	student["Josh"] = 31
	fmt.Println(student["Josh"])
	fmt.Println(len(student))
	fmt.Println(student)
}
