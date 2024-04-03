package main

import "fmt"

func main() {

	var names = []string{"Josh", "Alex", "Raechel", "Nick", "Jonathan"}

	for _, value := range names {
		fmt.Println(value)
	}
	fmt.Println("The gang's all here!")
}
