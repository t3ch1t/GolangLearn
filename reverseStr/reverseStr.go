package main

import  (
        "fmt"
)   

func main () {

    var testStr string = "HelloWorld"
    var newStr string = ""

    for _, i := range testStr {
        newStr = string(i) + newStr
        fmt.Println(i)
    }
    
    fmt.Println(newStr)

    return
}

