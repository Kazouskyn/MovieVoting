package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Movie Struct (Model)
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
