package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Movie Struct (Model)
type Movie struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Ranking []int  `json:"ranking"`
}

//movies will be a slice of article
var movies []Movie

//Data is createed for holding the key pairs of movies and their average voting totals
var Data map[string][]int

func main() {

	fmt.Println("REST API For Voting On Movies")

	//Mock Data
	slice := [...]int{1, 2, 3, 4, 5}
	movies = append(movies, Movie{ID: "1", Title: "Donnie Darko", Ranking: slice[1:4]})
	movies = append(movies, Movie{ID: "2", Title: "Halloween", Ranking: slice[0:3]})
	movies = append(movies, Movie{ID: "3", Title: "Interstellar", Ranking: slice[2:5]})
	//		Movie{MovieTitle: "Tombstone", Description: "Movie Description 1", VotingValues: slice[0:3]},
	//		Movie{MovieTitle: "Adventureland", Description: "Movie Description 2", VotingValues: slice[0:3]},
	//}
	handleRequests()
}

//handleRequests is a route handler that will handle all request to restAPI
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api", postMovies).Methods("POST")
	myRouter.HandleFunc("/api", getMovies).Methods("GET")
	myRouter.HandleFunc("/api/results", resultsPage).Methods("POST")
	myRouter.HandleFunc("/api/{ID}", getMovie).Methods("GET")       //can't get to work????
	myRouter.HandleFunc("/api/{ID}", deleteMovie).Methods("DELETE") //cant get to work????
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

//the homePage is the homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage \n")
	fmt.Println("Endpoint Hit: homePage")
}

// Get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getMovies endpint")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

// Post movies to /api page
func postMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to api postMovie page \n")
	fmt.Println("Endpoint hit: postMovies endpint")
	w.Header().Set("Content-Type", "application/json")
	var mover Movie
	_ = json.NewDecoder(r.Body).Decode(&mover)
	mover.ID = strconv.Itoa(rand.Intn(100000000)) //Mock ID
	movies = append(movies, mover)
	json.NewEncoder(w).Encode(mover)
}

//the resultsPage will display the movie title along with the total movie votes
func resultsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to results page \n")
	fmt.Println("Endpoint hit: resultsPage endpint")
	totalMovieRankings(movies)
	json.NewEncoder(w).Encode(Data)
}

// Total the rankings of the movies
func totalMovieRankings(movies []Movie) {
	Data = make(map[string][]int)

	for _, mov := range movies {
		name := mov.Title
		vals := mov.Ranking
		total := 0
		for _, num := range vals {
			total = total + num
		}
		Data[name] = append(Data[name], total)
	}
}

//WILL NOT WORK FOR SOME REASON
// Get single movie //WILL NOT WORK FOR SOME REASON
func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getMovie endpint")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range movies {
		if item.ID == params {
			json.NewEncoder(w).Encode(item)
			//return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

//WILL NOT WORK FOR SOME REASON
// Delete movie //WILL NOT WORK FOR SOME REASON
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteMovie endpint")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
