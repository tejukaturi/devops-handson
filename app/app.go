package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello DevOps Version 3")
}

unc main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}
