package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Movie Struct (Model)
type Movie struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Ranking string `json:"rating"`
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

// will check if a movie with the title is already in the DB
func isMovieInDB(title string) bool {
	//fmt.Println("in is Movie in DB")

	var records []Movie
	// gets all records in DB and stores them in records
	DB.Find(&records)
	// itterate over all records to find if movie is in DB
	// TO DO see if having the updateMovieRating() in side of the if to be usefull
	// would eliminate haveing to search through all of the DB twice
	for _, hold := range records {
		if hold.Title == title {
			//fmt.Println("the move:", title, " exists in database")
			return true
		}
	}
	return false
}

// will update a moive ranking with an added value for the movie incoming
func updateMovieRating(newRanking string, title string) {

	//fmt.Println("updateMovieRating hit")
	// convert the string to an int so it can be added to the old ranking
	newIntRanking, err := strconv.Atoi(newRanking)
	if err != nil {
		fmt.Printf("something went wrong with casting to int error: %v\n", err)
		os.Exit(1)
	}

	var records []Movie
	DB.Find(&records)

	for _, mov := range records {
		if mov.Title == title {
			oldrank, err := strconv.Atoi(mov.Ranking)
			if err != nil {
				fmt.Printf("something went wrong with casting to int error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("old rank:", oldrank)
			// once old movie is found add newRanking to ranking
			oldrank = oldrank + newIntRanking
			updatedRanking := strconv.Itoa(oldrank)
			mov.Ranking = updatedRanking
			//fmt.Println("new rank:", mov.Ranking)
			DB.Save(mov)
		}
	}
}
