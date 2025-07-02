package db

import "Movie/model"

var topMovies []model.TopMovies
var users []model.User

func init() {
	topMovies = []model.TopMovies{
		{Filmid: 1, FilmName: "The Shawshank Redemption", AverageRating: 9.3, TotalVotes: 2345678},
		{Filmid: 2, FilmName: "The Godfather", AverageRating: 9.2, TotalVotes: 2345678},
		{Filmid: 3, FilmName: "The Dark Knight", AverageRating: 9.0, TotalVotes: 2345678},
	}
	users = []model.User{
		{ID: 1, Username: "Talberto", Password: "Password123"},
		{ID: 2, Username: "JaneDoe", Password: "SecurePassword"},
	}
}

func GetTopMovies() []model.TopMovies {
	return topMovies
}

func GetUsers() []model.User {
	return users
}
