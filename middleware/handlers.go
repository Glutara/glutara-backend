package middleware

import (
	"net/http"
	"fmt"
)

// TODO :
// Response format

// Create a global instance of Repo
// SetRepo sets the instance of Repo

// MainHandler for dummy test
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "Hello, Welcome to Glutara Web Service!")
}

// MainHandler for dummy test
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "About Glutara")
}