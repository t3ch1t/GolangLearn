package main

import (
    "fmt"
    "io/fs"
    "flag"
    "os"
)

func main() {
    filepath := flag.String("file", "$HOME/.bashrc", "Path to fetch file data")  
    flag.Parse()
    
    isValidPath := fs.ValidPath(*filepath)
    if isValidPath != true {
        fmt.Println(isValidPath)
        fmt.Println("That is not a valid filepath. Now exiting.")
        os.Exit(2)
    }

    fileInfo, err := os.Stat(*filepath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(fileInfo)
}
