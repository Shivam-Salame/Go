package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	// running server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hello from greeter")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from serverHome")

	w.Write([]byte("<h1>Welcome to serverHome response writer</h1>"))
}
