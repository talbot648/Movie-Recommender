package service

import (
	"Movie/db"
	"Movie/db/postgres"
	"errors"
	"fmt"
)

func GetTopMovies() ([]db.TopMovies, error) {
	TopMovies, err := postgres.GetTopMovies()

	if err != nil {
		fmt.Println("Error getting top movies from DB", err)
		return nil, errors.New("There was an error getting the users from the database")
	}
	return TopMovies, nil
}

func GetMovieDetails(id int) (db.MovieDetails, error) {
	movieDetails, err := postgres.GetMovieDetails(id)

	if err != nil {
		fmt.Println("Error getting Movie details from DB", err)
		return movieDetails, errors.New("There was an error getting the movie details")
	}
	return movieDetails, nil
}
