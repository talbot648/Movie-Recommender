package db

type TopMovies struct {
	FilmName      string
	AverageRating float64
	TotalVotes    int
}

var topMovies []TopMovies

func init() {
	topMovies = []TopMovies{
		{FilmName: "The Shawshank Redemption", AverageRating: 9.3, TotalVotes: 2345678},
		{FilmName: "The Godfather", AverageRating: 9.2, TotalVotes: 2345678},
		{FilmName: "The Dark Knight", AverageRating: 9.0, TotalVotes: 2345678},
	}
}

func GetTopMovies() []TopMovies {
	return topMovies
}
