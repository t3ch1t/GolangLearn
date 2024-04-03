package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userHome)
}
