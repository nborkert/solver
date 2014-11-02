package webhandler

import (
    "fmt"
    "net/http"
//    "simple"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Solver")
//    fmt.Fprint(w, simple.Echo())
}
