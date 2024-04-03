package main

import "fmt"

func main() {

	evenNum := [5]int{0, 2, 4, 6, 8}

	fmt.Println(evenNum)

	for i, value := range evenNum {

		fmt.Println(value, i)
	}

	numSlice := []int{5, 4, 3, 2, 1}

	sliced := numSlice[3:5]
	fmt.Println(numSlice)
	fmt.Println(sliced)

}
