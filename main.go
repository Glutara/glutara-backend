package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Welcome to Glutara Web Service!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About Page")
}

func main() {
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/", mainHandler)
    r.HandleFunc("/about", aboutHandler)

    // Start the HTTP server on port 8080 using the router
    fmt.Println("Server is running on :8080...")
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
