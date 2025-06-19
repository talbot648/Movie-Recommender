package main

import (
	// Assuming db is a package that contains the TopMovies struct and GetTopMovies function
	"Movie/db/postgres" // Assuming postgres is a package that initializes the database connection
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	connectionString, err := postgres.GetDBConnectionString("../scripts/config.ini") // Adjust the path to your config file as needed
	if err != nil {
		fmt.Println("Error getting DB connection string:", err)
		return
	}

	if err := postgres.InitDB(connectionString); err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer postgres.DB.Close()

	router := http.NewServeMux()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/api/topMovies", getTopMovies)
	router.HandleFunc("/api/movieDetails", getMovieDetails)

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", CorsMiddleware(router))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		// Continue with the next handler
		next.ServeHTTP(writer, request)
	})
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func getTopMovies(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Printf("got /api/topMovies request\n")

	topMovies, errMovie := postgres.GetTopMovies()
	if errMovie != nil {
		http.Error(writer, "Error getting movies", http.StatusInternalServerError)
		return
	}

	topMoviesJSON, errMarshal := json.Marshal(topMovies)
	if errMarshal != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return

	}

	writer.Header().Set("Content-Type", "application/json")

	_, err := writer.Write([]byte(topMoviesJSON))
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func getMovieDetails(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ids := request.URL.Query()["id"]
	id, _ := strconv.Atoi(ids[0])
	fmt.Printf("got /api/movieDetails request for id=%d\n", id)

	movieDetails, errMovie := postgres.GetMovieDetails(id)
	if errMovie != nil {
		http.Error(writer, "Error getting movie details", http.StatusInternalServerError)
		fmt.Print(errMovie)
		return
	}

	movieDetailsJSON, errMarshal := json.Marshal(movieDetails)
	if errMarshal != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return

	}
	writer.Header().Set("Content-Type", "application/json")
	_, err := writer.Write([]byte(movieDetailsJSON))
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
