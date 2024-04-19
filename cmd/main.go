package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Vercel and Go!")
	})

	port := ":9000"
	fmt.Printf("Listening on %s...\n", port)
	http.ListenAndServe(port, nil)
}
