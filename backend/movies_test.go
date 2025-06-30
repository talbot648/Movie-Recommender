package main

import (
	"Movie/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetTopMoviesHandler(t *testing.T) {
	//Arrange
	req, err := http.NewRequest("GET", "/api/topMovies", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a handler
	handler := http.HandlerFunc(getTopMovies)

	want := []model.TopMovies{
		{
			Filmid:        1,
			FilmName:      "The Shawshank Redemption",
			AverageRating: 9.3,
			TotalVotes:    2345678,
		},
		{
			Filmid:        2,
			FilmName:      "The Godfather",
			AverageRating: 9.2,
			TotalVotes:    2345678,
		},
		{
			Filmid:        3,
			FilmName:      "The Dark Knight",
			AverageRating: 9.0,
			TotalVotes:    2345678,
		},
	}

	expectedJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("Failed to marshal expected JSON: %v", err)
	}
	//Act
	handler.ServeHTTP(rr, req)

	//Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != string(expectedJSON) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), want)
	}

	var got []model.TopMovies

	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("handler returned unexpected body: got %v, expected %v", got, want)
	}
}
