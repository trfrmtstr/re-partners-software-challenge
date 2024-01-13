package main

import (
	"fmt"
	"net/http"
	"re-partners-take-home-assignment/handlers"
)

// Sets up handlers for algorithm API call and serving index.html then runs web server
func main() {
	http.HandleFunc("/find-optimal-packing", handlers.FindOptimalPackingHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Optimal packing web server is running...")
	http.ListenAndServe(":8080", nil)
}
