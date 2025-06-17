package main

import (
	"Movie/db"          // Assuming db is a package that contains the TopMovies struct and GetTopMovies function
	"Movie/db/postgres" // Assuming postgres is a package that initializes the database connection
	"encoding/json"
	"fmt"
	"net/http"
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

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/topMovies", getTopMovies)

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

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

	topMovies := db.GetTopMovies() // Assuming db.GetTopMovies() returns a slice of TopMovies

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
