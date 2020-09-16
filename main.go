package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := connectDatabase()
	if err != nil {
		fmt.Printf("There was an issue connecting to the database error: %v\n", err)
		os.Exit(1)
	}

	router.Use(cors.New(cors.Config{
		AllowCredentials:       false,
		AllowAllOrigins:        true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		//AllowOrigins:     []string{"*"},
		AllowMethods:  []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE", "*"},
		AllowHeaders:  []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "*"},
		ExposeHeaders: []string{"Content-Length", "*"},

		//AllowOriginFunc: func(origin string) bool {
		//	return true
		//},
	}))

	// Routes
	router.GET("/api", findMovies)
	router.GET("/api/:id", findMovie)
	//router.GET("/api/:title", findMovie)
	router.POST("/api", createMovie)
	router.PATCH("/api/:id", updateMovie)
	router.DELETE("/api/:id", deleteMovie)
	router.DELETE("/", clearDataBase)

	// Run the server on port 3000
	router.Run(":3000")
}
