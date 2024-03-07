package main

import (
	"fmt"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query parameter "text" from the request
	text := r.URL.Query().Get("text")

	// If the text parameter is not provided, respond with an error
	if text == "" {
		http.Error(w, "Please provide text parameter", http.StatusBadRequest)
		return
	}

	// Intentional error: Not checking for errors when using fmt.Fprintf
	_, err := fmt.Fprintf(w, "Echo: %s", text)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// Register the echoHandler function to handle requests to the "/echo" path
	http.HandleFunc("/echo", echoHandler)

	// Intentional error: Not checking for errors when starting the server
	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
