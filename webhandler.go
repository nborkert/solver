package solver

import (
    "fmt"
    "net/http"
    "runtime"
//    "simple"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Fprint(w, "Solver")
    fmt.Fprint(w, runtime.NumCPU())
//    fmt.Fprint(w, simple.Echo())
}
