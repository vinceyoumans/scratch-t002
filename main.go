package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "<p>Hello from Go!</p>")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write response: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Failed to write response: %v\n", err)
			if err != nil {
				return
			}
		}
		os.Exit(1)
	}
}
