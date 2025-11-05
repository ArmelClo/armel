package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	fmt.Println("Server listing to :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello de Armel")
	if err != nil {
		return
	}
}
