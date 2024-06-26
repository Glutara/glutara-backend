package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"glutara/router"
)

func init() {
	// Load environment variables from
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Configure router to include routes and middlewares
	router.ConfigureRouter(r)

    // Start server
	// Use `PORT` provided in environment or default to 8080
  	r.Run(envPortOr("8080"))
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