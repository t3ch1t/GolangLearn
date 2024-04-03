package main

import "fmt"

func main() {
	var str1 string = "hello"
	var str2 string = "goodbye"
	var str3 string = ""

	for i := 0; i < len(str1) || i < len(str2); i++ {

		if i < len(str1) {
			str3 += string(str1[i])
		}

		if i < len(str2) {
			str3 += string(str2[i])
		}

	}

	fmt.Println(str3)

}
