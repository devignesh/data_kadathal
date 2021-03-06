package main

import (
    "fmt"
    "net/http"
)

type Mymux struct {
}

func (p *Mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayhelloName(w, r)
        return
    }
    http.NotFound(w, r)
    return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello vicky")
    fmt.Println(sayhelloName)
}

func main() {
    mux := &Mymux{}
    http.ListenAndServe(":8080", mux)
}
