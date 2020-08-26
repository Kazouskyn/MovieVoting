package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createMoiveInput is a struct for createMovie
type createMoiveInput struct {
	Title   string //`json:"title" binding:"required"`
	Ranking string //`json:"ranking" binding:"required"`
}

// updateMovieInput is a struct for updateMovie
type updateMovieInput struct {
	ID      int    //`json:"id"`
	Title   string //`json:"title"`
	Ranking string //`json:"ranking"`
}

// findMovies will find and return all movies
// GET /api
func findMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var movies []Movie
	DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// findMovie will find a specified movie based off of the id
// GET /api/:id
func findMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Get model if exist
	var movie Movie
	if err := DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// createMovie will create a new movie and add it to the data base
// POST /api
func createMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Validate input
	var input createMoiveInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create movie
	movie := Movie{Title: input.Title, Ranking: input.Ranking}
	DB.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// updateMovie will update a movie based off of the id
// PATCH /api/:id
func updateMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Get model if exist
	var movie Movie
	if err := DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input updateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&movie).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// deleteMovie will delete a movie based off of an id
// DELETE /api/:id
func deleteMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Get model if exist
	var movie Movie
	if err := DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func clearDataBase(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	fmt.Println("you are going to delete the entire database are you sure? (y/n)")
	//reader := bufio.NewReader(os.Stdin)
	//input, err := reader.ReadString('\n')
	//if err != nil {
	//fmt.Printf("something went wrong reading input error: %v\n", err)
	//}
	var movie Movie
	DB.Unscoped().Delete(&movie)
}
