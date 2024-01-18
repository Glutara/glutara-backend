package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"glutara/router"
)

func main() {
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	
	// Activate router
	r := router.Router()

    // Start server
	// Use `PORT` provided in environment or default to 8080
  	var port = envPortOr("8080")
	fmt.Println("Starting server...")
	fmt.Println("Listening from http://localhost" + port + "/api")
	log.Fatal(http.ListenAndServe(port, r))
}

// PORT handling
// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":8080"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}