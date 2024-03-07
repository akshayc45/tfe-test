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

	// Echo back the received text in the response
	fmt.Fprintf(w, "Echo: %s", text)
}

func main() {
	// Register the echoHandler function to handle requests to the "/echo" path
	http.HandleFunc("/echo", echoHandler)

	// Start the server and listen on port 8080
	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
