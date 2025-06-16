package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, World!")
	})

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
