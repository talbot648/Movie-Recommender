package main

import (
	"Movie/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetUsersHandler(t *testing.T) {
	// ARRANGE
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a handler
	handler := http.HandlerFunc(getUsers)

	expected := []model.User{
		{
			ID:       1,
			Username: "Talberto",
			Password: "Password123",
		},
		{

			ID:       2,
			Username: "JaneDoe",
			Password: "SecurePassword",
		},
	}

	//ACT
	handler.ServeHTTP(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual []model.User
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateUserHandler(t *testing.T) {
	// ARRANGE
	// Create a new HTTP request with a JSON body
	expected := model.User{
		ID:       3,
		Username: "NewUser",
		Password: "NewPassword",
	}
	body, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	var actual model.User
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if expected.ID != actual.ID || expected.Username != actual.Username || expected.Password != actual.Password {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}
