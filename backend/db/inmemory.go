package db

import "Movie/model"

var topMovies []model.TopMovies

func init() {
	topMovies = []model.TopMovies{
		{Filmid: 1, FilmName: "The Shawshank Redemption", AverageRating: 9.3, TotalVotes: 2345678},
		{Filmid: 2, FilmName: "The Godfather", AverageRating: 9.2, TotalVotes: 2345678},
		{Filmid: 3, FilmName: "The Dark Knight", AverageRating: 9.0, TotalVotes: 2345678},
	}
}

func GetTopMovies() []model.TopMovies {
	return topMovies
}
