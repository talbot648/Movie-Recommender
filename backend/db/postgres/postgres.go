package postgres

import (
	"Movie/db" // Assuming db is the package where TopMovies struct is defined
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gopkg.in/ini.v1"
)

var DB *sql.DB

func InitDB(connectionString string) error {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	DB = db

	// Check if the connection is established
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}
	fmt.Println("Successfully connected to the Database!")
	return nil
}

func GetDBConnectionString(path string) (string, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return "", fmt.Errorf("failed to load config file: %w", err)
	}
	section := cfg.Section("postgres")
	host := section.Key("host").String()
	port := section.Key("port").String()
	user := section.Key("username").String()
	password := section.Key("password").String()
	database := section.Key("db").String()
	sslmode := section.Key("sslmode").String()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, database, sslmode)

	return connectionString, nil
}

func GetTopMovies() ([]db.TopMovies, error) {
	rows, err := DB.Query("SELECT film_id, film_name, average_rating, total_votes FROM top_films_new")

	if err != nil {
		fmt.Println("failed to query the databse", err)
		return []db.TopMovies{}, errors.New("Database could not be queried")
	}
	defer rows.Close()

	topMovies := []db.TopMovies{}

	for rows.Next() {
		var movie db.TopMovies
		if err := rows.Scan(&movie.Filmid, &movie.FilmName, &movie.AverageRating, &movie.TotalVotes); err != nil {
			fmt.Println("failed to scan row", err)
			return []db.TopMovies{}, errors.New("Database could not be queried")
		}
		topMovies = append(topMovies, movie)
	}
	return topMovies, nil
}

func GetMovieDetails(id int) (db.MovieDetails, error) {
	var movie db.MovieDetails
	row := DB.QueryRow(`SELECT id, title, Adult, genres, language, overview, releasedate, voteaverage, votecount
      FROM cl.movies_metadata
      WHERE id = $1`, id)

	// Scan into your struct fields
	err := row.Scan(
		&movie.Filmid,
		&movie.FilmName,
		&movie.Adult,
		&movie.Genres,
		&movie.Language,
		&movie.Overview,
		&movie.ReleaseDate,
		&movie.AverageRating,
		&movie.TotalVotes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return movie, fmt.Errorf("no movie found with id %d", id)
		}
		return movie, fmt.Errorf("query error: %w", err)
	}

	return movie, nil
}
