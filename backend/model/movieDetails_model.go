package model

import "time"

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
