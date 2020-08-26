package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.Use(static.Serve("/", static.LocalFile("./views", true)))

	err := connectDatabase()
	if err != nil {
		fmt.Printf("There was an issue connecting to the database error: %v\n", err)
		os.Exit(1)
	}

	//router.Use(apiMiddleware(DB))

	// Routes
	router.GET("/api", findMovies)
	router.GET("/api/:id", findMovie)
	router.POST("/api", createMovie)
	router.PATCH("/api/:id", updateMovie)
	router.DELETE("/api/:id", deleteMovie)
	router.DELETE("/", clearDataBase)

	// Run the server on port 8080
	router.Run(":3000")
}

/*//Movie Struct (Model)
type Movie struct {
	ID      int    //`json:"id"`
	Title   string //`json:"title"`
	Ranking string //`json:"ranking"`
}

// DB is the data base
var DB *gorm.DB

// connectDatabase will connect to the database
func connectDatabase() error {
	database, err := gorm.Open("sqlite3", "movie.db")
	if err != nil {
		fmt.Printf("There was an issue connecting to the database error: %v\n", err)
		os.Exit(1)
	}

	database.AutoMigrate(&Movie{})

	DB = database
	return err
}

func apiMiddleware(DB gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", DB)
		c.Next()
	}
}*/
