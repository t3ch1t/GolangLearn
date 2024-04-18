package main

import ("fmt"
        "os"
        "log"
        "os/user"
    )

func main () {

    hostname, err := os.Hostname()
    if err != nil {
        log.Fatal("There has been an error geting the hostname", err)
    }
    currentUser, err := user.Current() 
   if err != nil {
    log.Fatal("There has been an error", err)
    }
    groups, err := currentUser.GroupIds()
   if err != nil {
        log.Fatal("There has been an error", err)
    }
    fmt.Println("Hostname: ", hostname)
    fmt.Println("Username: ", currentUser.Username)
    fmt.Println("Name: ", currentUser.Name)
    fmt.Println("Uid: ", currentUser.Uid)
    fmt.Println("Gid: ", currentUser.Gid)
    fmt.Println("Groups: ", groups)
    fmt.Println("Home Directory: ", currentUser.HomeDir)
    
}
