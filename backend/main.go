package main

import (
	"Movie/api"
	"Movie/db"
	"Movie/db/postgres"
	"Movie/model"
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
	router.HandleFunc("GET /api/topMovies", api.GetTopMovies)
	router.HandleFunc("GET /api/movieDetails/{id}", api.GetMovieDetails)
	router.HandleFunc("/api/users", handleUsers)

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

func handleUsers(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var user model.User
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			fmt.Println("Error decoding user:", err)
			http.Error(writer, "Bad request", http.StatusBadRequest)
			return
		}
		id := db.AddUser(user)
		user.ID = id
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(user)

		return
	}
	fmt.Printf("got /api/users request\n")
	users := db.GetUsers()

	json.NewEncoder(writer).Encode(users)

}
