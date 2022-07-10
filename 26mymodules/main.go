package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Wlcome to mymodules")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	http.ListenAndServe(":4000", r)

}

func greeter() {
	fmt.Println("Welcome to my modules")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my modules<h1>"))
}
