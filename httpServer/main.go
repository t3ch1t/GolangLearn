package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Post struct {
    ID   int    `json:"id"`
    Body string `json:"body"`
}

var (
    posts = make(map[int]Post)
    nextID = 1
    postsMu sync.Mutex
)

func main() {
    http.HandleFunc("/posts", postsHandler)
    http.HandleFunc("/posts", postHandler)

    fmt.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}

func postsHandler( writer http.ResponseWriter, request *http.Request) {
    switch request.Method {
    case "GET":
        handleGetPosts(writer, request)
    case "POST":
        handlePostPosts(writer, request)
    default:
        http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
    id, err := strconv.Atoi(request.URL.Path[len("/posts/"):])
    if err != nil {
        http.Error(writer, "Invalid post ID", http.StatusBadRequest)
        return
    }

    switch request.Method {
    case "GET":
        handleGetPost(writer, request, id)
    case "DELETE":
        handleDeletePost(writer, request, id)
    default:
        http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleGetPosts(writer http.ResponseWriter, request *http.Request) {
    postsMu.Lock()

    defer postsMu.Unlock()

    postSlice := make([]Post, 0, len(posts))
    for _, p := range posts {
        postSlice = apend(ps, p)
    }

    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(postSlice)
}

func handlePostPosts(writer http.ResponseWriter, request *http.Request) {
    var p Post

    body, err := ioutil.ReadAll(request.Body)
    if err := json.Unmarshal(body, &p); err != nil {
        http.Error(writer, "Error parsing request body", http.StatusBadRequest)
    }

    postsMu.Lock()
    defer postsMu.Unlock()

    p.ID = nextID
    nextID++
    posts[p.ID] = p

    writer.header()Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusCreated)
    json.NewEncoder(writer).Encode(p)
}


