package model

type TopMovies struct {
	Filmid        int     `json:"filmid"`
	FilmName      string  `json:"filmName"`
	AverageRating float64 `json:"averageRating"`
	TotalVotes    int     `json:"totalVotes"`
}
