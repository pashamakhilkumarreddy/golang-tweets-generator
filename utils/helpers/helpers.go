package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	m "github.com/pashamakhilkumarreddy/golang-tweets-generator/models"
)

func ParseReviewsData(filePath string) ([]m.Review, error) {
	file, err := os.Open(filePath)
	if err != nil {
		LogError(err)
		return nil, errors.New("invalid file")
		// os.Exit(1)
	}
	defer file.Close()
	bytesData, err := ioutil.ReadAll(file)
	var reviews []m.Review
	if err != nil {
		LogError(err)
		return nil, errors.New("invalid file")
	}
	json.Unmarshal(bytesData, &reviews)
	return reviews, nil
}

func ParseMoviesData(filePath string) ([]m.Movie, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Invalid file")
	}
	bytesData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Invalid file")
	}
	var movies []m.Movie
	json.Unmarshal(bytesData, &movies)
	defer file.Close()
	return movies, nil
}

func ParseJsonData(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Invalid file")
	}
	bytesData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Invalid file")
	}
	var result []map[string]interface{} // result[0]["title"]
	json.Unmarshal(bytesData, &result)
	defer file.Close()
	return result, nil
}

func GetYearFromTitle(movies []m.Movie, title string) int {
	for _, movie := range movies {
		if movie.Title == title {
			return movie.Year
		}
	}
	return 0
}

func CreateTweet(title, review, year string, score uint16) string {
	const (
		star = "*"
		half = "½"
	)
	intVal, frac := math.Modf(float64(score) / 20.0)
	formattedScore := strings.Repeat(star, int(intVal))
	formattedYear := year
	if year == "0" {
		formattedYear = ""
	} else {
		formattedYear = fmt.Sprintf(" (%s)", year)
	}
	if frac > 0.5 {
		formattedScore += star
	} else if intVal < 5.0 && frac <= 0.5 {
		formattedScore += half
	}
	tweet := fmt.Sprintf("%s%v: %s %v", title, formattedYear, review, formattedScore)
	if len(tweet) > 140 {
		truncTitle := title
		truncRev := review
		if len(title) > 25 {
			truncTitle = strings.Trim(title[:25], "")
		}
		if len(tweet) > 140 && len(review) > 25 {
			truncRev = strings.Trim(review[:(len(review)-(len(tweet)-140))], "")
		}
		tweet = fmt.Sprintf("%s%v: %s %v", truncTitle, formattedYear, truncRev, formattedScore)
	}
	return tweet
}

func LogError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
