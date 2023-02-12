package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func main() {
	r := gin.Default()

	// Read the JSON file that defines the routes
	data, err := os.ReadFile("routes.json")
	if err != nil {
		log.Fatalf("Failed to read routes file: %v", err)
	}

	// Parse the JSON data into a slice of Route structures
	var routes []Route
	err = json.Unmarshal(data, &routes)
	if err != nil {
		log.Fatalf("Failed to parse routes file: %v", err)
	}

	// Map of handlers for each HTTP method
	handler := map[string]gin.HandlerFunc{
		"GET": func(c *gin.Context) {
			c.String(200, "Handling GET request")
		},
		"POST": func(c *gin.Context) {
			c.String(200, "Handling POST request")
		},
		// Add additional cases for other HTTP methods as needed
	}

	// Loop over the routes and add them to the Gin router
	for _, route := range routes {
		if fn, exists := handler[route.Method]; exists {
			r.Handle(route.Method, route.Path, fn)
		} else {
			log.Fatalf("Unsupported HTTP method: %s", route.Method)
		}
	}

	err = r.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
