package db

import "time"

type TopMovies struct {
	Filmid        int
	FilmName      string
	AverageRating float64
	TotalVotes    int
}

type MovieDetails struct {
	Filmid        int
	FilmName      string
	Adult         bool
	Genres        string
	Language      string
	Overview      string
	ReleaseDate   time.Time
	AverageRating float64
	TotalVotes    float64
}

var topMovies []TopMovies

func init() {
	topMovies = []TopMovies{
		{Filmid: 1, FilmName: "The Shawshank Redemption", AverageRating: 9.3, TotalVotes: 2345678},
		{Filmid: 2, FilmName: "The Godfather", AverageRating: 9.2, TotalVotes: 2345678},
		{Filmid: 3, FilmName: "The Dark Knight", AverageRating: 9.0, TotalVotes: 2345678},
	}
}

func GetTopMovies() []TopMovies {
	return topMovies
}
