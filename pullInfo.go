package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func sdf() {
	fmt.Println("Endpoint hit: getData endpint")

	url := "http://localhost:8080/api"
	res, err := http.Get(url)
	fmt.Println("error:", err)
	if err != nil {
		fmt.Printf("something went wrong with the Get request error: %v\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("something went wrong with the reading the body error: %v\n", err)
		os.Exit(1)
	}

	var data Movie

	// unmarshall
	json.Unmarshal(body, &data)

	fmt.Printf("Results: %v\n", data)
	fmt.Println("data.votingval", data.VotingValues)
	fmt.Println("data.Description", data.Description)
	fmt.Println("data.VotingValues", data.VotingValues)

}
