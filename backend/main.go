package main

import (
	"Movie/api"
	"Movie/db"
	"Movie/db/postgres"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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

func main() {

	connectionString, err := postgres.GetDBConnectionString("../scripts/config.ini")
	if err != nil {
		log.Fatal("Error getting connection string")
		return
	}
	if err := postgres.InitDB(connectionString); err != nil {
		fmt.Println("Encountered Error initializing database:", err)
		return
	}
	defer postgres.DB.Close()

	router := http.NewServeMux()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/api/topMovies", api.GetTopMovies)
	router.HandleFunc("/api/movieDetails/{id}", api.GetMovieDetails)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server listening on port %s\n", port)
	err = http.ListenAndServe(":"+port, CorsMiddleware(router))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func getTopMovies(writer http.ResponseWriter, request *http.Request) {
	topMovies := db.GetTopMovies()

	topMoviesJSON, errMarshal := json.Marshal(topMovies)
	if errMarshal != nil {
		// Handle error if marshalling fails
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err := writer.Write([]byte(topMoviesJSON))
	if err != nil {
		// Handle error if writing response fails
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
