package main

import (
    "fmt"
    "log"
    "net/http"
    )

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "A journey of a thousand miles begins with a single step. - Lao Tzu")
}

func main() {
    http.HandleFunc("/",handler)
    log.Fatal(http.ListenAndServe(":8080",nil))
}

