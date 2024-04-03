package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is only a test.")
	file.Close()

	stream, err := os.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)

	fmt.Println(s1)

}
