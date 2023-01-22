package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fmt.Print("Starting web server at http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
