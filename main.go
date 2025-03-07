package main

import (
	"fmt"
	"net/http"
	"os"
	su "t002/setup"

	"t002/internals"
)

func main() {
	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir(su.UrlToTemplate + "/static"))

	http.Handle("/", fs)

	// Define routes for the application
	http.HandleFunc("/login", internals.Login)
	http.HandleFunc("/authenticate", internals.Authenticate)
	http.Handle("/protected", internals.Authorize(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/landing-auth", http.StatusSeeOther)
	})))
	http.HandleFunc("/landing-auth", internals.LandingAuth)
	http.HandleFunc("/landing", internals.Landing)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<p>Hello from Go!</p>")
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		os.Exit(1)
	}
	
}
