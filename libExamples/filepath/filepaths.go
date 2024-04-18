package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    path := "/home/josh/Documents/Code/Go/quiz"
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        fmt.Println(path)
        return nil
    })
}
