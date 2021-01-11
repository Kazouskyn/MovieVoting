package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// createMoiveInput is a struct for createMovie
type createMoiveInput struct {
	ID      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Ranking string `json:"rating" binding:"required"`
}

// updateMovieInput is a struct for updateMovie
type updateMovieInput struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Ranking string `json:"rating"`
}

// findMovies will find and return all movies
// GET /api
func findMovies(c *gin.Context) {
	var movies []Movie
	DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// findMovie will find a specified movie based off of the id
// GET /api/:id
func findMovie(c *gin.Context) {
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
	data, err := c.GetRawData()
	if err != nil {
		fmt.Printf("something went wrong with getting raw data error: %v\n", err)
		os.Exit(1)
	}
	var smhover Movie
	var movie Movie
	var trimmed string
	slice := strings.SplitAfter(string(data), "}")
	// will loop over all of the slices
	for counter, test := range slice {
		//this accounts for the end string that has no valuable data
		if test == "]" {
			break
		}
		//this accounts for the first string so the "[" that is fed to me is taken off of the string
		if counter == 0 {
			// trim the "[" off so the data is in correct format
			trimmed = strings.TrimPrefix(test, "[")
			err = json.Unmarshal([]byte(trimmed), &smhover)
			if err != nil {
				fmt.Printf("something went wrong with Unmarshalling the data error: %v\n", err)
				os.Exit(1)
			}
			// check if the movie is in the DB
			// if the movie is in the DB than add the incoming ranking to the total
			ok := isMovieInDB(smhover.Title)
			if ok {
				updateMovieRating(smhover.Ranking, smhover.Title)
			} else {
				// Create movie
				movie = Movie{Title: smhover.Title, Ranking: smhover.Ranking}
				DB.Create(&movie)
			}
		} else {
			trimmed = strings.TrimPrefix(test, ",")
			err = json.Unmarshal([]byte(trimmed), &smhover)
			if err != nil {
				fmt.Printf("something went wrong with Unmarshalling the data error 2: %v\n", err)
				os.Exit(1)
			}
			// check if the movie is in the DB
			// if the movie is in the DB than add the incoming ranking to the total
			ok := isMovieInDB(smhover.Title)
			if ok {
				updateMovieRating(smhover.Ranking, smhover.Title)
			} else {
				// Create movie
				movie = Movie{Title: smhover.Title, Ranking: smhover.Ranking}
				DB.Create(&movie)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})

}

// updateMovie will update a movie based off of the id
// PATCH /api/:id
func updateMovie(c *gin.Context) {
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
	var movie Movie
	DB.Unscoped().Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
