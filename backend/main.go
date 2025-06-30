package main

import (
	"Movie/api"
	"Movie/db/postgres"
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
