package solver

import (
	"fmt"
	"net/http"
	"runtime"
	"io/ioutil"
	//    "simple"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/players", playersHandler)
	http.HandleFunc("/roster", rosterHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Fprint(w, "Solver")
	fmt.Fprint(w, runtime.NumCPU())
	//    fmt.Fprint(w, simple.Echo())
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := ioutil.ReadFile("json.txt")

	fmt.Fprintf(w, "%s", t)
}

func rosterHandler(w http.ResponseWriter, r * http.Request) {
	p, _ := ioutil.ReadFile("roster.txt")
	fmt.Fprintf(w, "%s", p)
}


