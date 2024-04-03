package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	const testStr string = "qw3erty7uiop0"
	var results []int

	for i := 0; i < len(testStr); i++ {
		slice := rune(testStr[i])

		if unicode.IsDigit(slice) == true {
			digit, err := strconv.Atoi(string(slice))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(digit)
			results = append(results, digit)
		}
	}

	fmt.Println(results)
}
