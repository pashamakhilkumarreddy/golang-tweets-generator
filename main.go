package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	h "github.com/pashamakhilkumarreddy/golang-tweets-generator/utils/helpers"
)

func main() {
	fmt.Println("Welcome to Golang Tweets Generator\n")
	args := os.Args[1:]
	if len(args) != 2 {
		h.LogError(errors.New("please input both reviews and movies json files"))
	}
	reviews, err := h.ParseReviewsData(args[0])
	if err != nil {
		h.LogError(err)
	}
	// fmt.Printf("Reviews are %v \n", reviews)

	movies, err := h.ParseMoviesData(args[1])
	if err != nil {
		h.LogError(err)
	}
	// fmt.Printf("\nMovies are %v\n", movies)

	for i := 0; i < len(reviews); i++ {
		var (
			title  = strings.Trim(reviews[i].Title, "")
			review = strings.Trim(reviews[i].Review, "")
			score  = reviews[i].Score
		)

		year := h.GetYearFromTitle(movies, title)
		tweet := h.CreateTweet(title, review, strconv.Itoa(year), score)
		fmt.Printf("Tweet is: %s\n", tweet)
	}
}
