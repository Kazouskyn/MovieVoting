package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Movie will hold information about a movie
type Movie struct {
	MovieTitle   string `json:"movieTitle"`
	Description  string `json:"description"`
	VotingValues []int  `json:"votingValues"`
}

//movies will be a slice of article
var movies []Movie

//Data is createed for holding the key pairs of movies and their average voting totals
var Data map[string][]float64

func main() {

	fmt.Println("REST API For Voting On Movies")
	slice := [...]int{1, 2, 3, 4, 5}
	movies = []Movie{
		Movie{MovieTitle: "Donnie Darko", Description: "Movie Description 1", VotingValues: slice[2:5]},
		Movie{MovieTitle: "Halloween", Description: "Movie Description 2", VotingValues: slice[1:4]},
		Movie{MovieTitle: "Tombstone", Description: "Movie Description 1", VotingValues: slice[0:3]},
		Movie{MovieTitle: "Adventureland", Description: "Movie Description 2", VotingValues: slice[0:3]},
	}
	handleRequests()
}

//handleRequests will handle all of the request to the 'server'
func handleRequests() {

	//myRouter := mux.NewRouter().StrictSlash(true)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api", returnAllMovies)
	//http.HandleFunc("/get", getData)
	http.HandleFunc("/results", resultsPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//the homePage is the homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

//returnAllMovies is where all of the movie data will be stored
func returnAllMovies(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: returnAllMovies endpint")
	json.NewEncoder(w).Encode(movies)
	averageMovies(movies)
}

//the resultsPage will display the movie title along with the average movie votes
func resultsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: resultsPage endpint")
	json.NewEncoder(w).Encode(Data)

}

/*
// Article - Our struct for all articles
type Article struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

*/
