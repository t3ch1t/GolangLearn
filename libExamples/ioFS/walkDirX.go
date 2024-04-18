package main


import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	root := "/home/josh/Documents/Code/Go"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}
