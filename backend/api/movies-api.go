package api

import (
	"Movie/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetTopMovies(writer http.ResponseWriter, request *http.Request) {

	topMovies, errMovie := service.GetTopMovies()
	if errMovie != nil {
		http.Error(writer, "Error getting movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(topMovies)

}

func GetMovieDetails(writer http.ResponseWriter, request *http.Request) {
	id, err := parseId(request.PathValue("id"))

	if err != nil {
		http.Error(writer, "Bad request ID", http.StatusBadRequest)
	}

	movieDetails, errMovie := service.GetMovieDetails(id)
	if errMovie != nil {
		http.Error(writer, "Error getting movie details", http.StatusInternalServerError)
		fmt.Print(errMovie)
		return
	}

	json.NewEncoder(writer).Encode(movieDetails)

}

func parseId(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error parisng ID:", err)
		return 0, err
	}
	return id, nil
}
