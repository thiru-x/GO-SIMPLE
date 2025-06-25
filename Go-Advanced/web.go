package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go server!")
	fmt.Fprintln(w, "HI DA  I AM THiRU GO SERVER!")

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
