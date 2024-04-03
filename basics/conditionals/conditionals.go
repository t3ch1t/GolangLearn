package main

import "fmt"

func main() {

	var i int = 1

	for i <= 5 {
		if i%2 == 0 {
			fmt.Println("even")
			i++
		} else {
			fmt.Println("odd")
			i++
		}
	}
}
