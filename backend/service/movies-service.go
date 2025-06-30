package service

import (
	"Movie/db/postgres"
	"Movie/model"
	"errors"
	"fmt"
)

func GetTopMovies() ([]model.TopMovies, error) {
	topMovies, err := postgres.GetTopMovies()

	if err != nil {
		fmt.Println("Error getting top movies from DB", err)
		return nil, errors.New("There was an error getting the users from the database")
	}
	return topMovies, nil
}

func GetMovieDetails(id int) (model.MovieDetails, error) {
	movieDetails, err := postgres.GetMovieDetails(id)

	if err != nil {
		fmt.Println("Error getting Movie details from DB", err)
		return movieDetails, errors.New("There was an error getting the movie details")
	}
	return movieDetails, nil
}
