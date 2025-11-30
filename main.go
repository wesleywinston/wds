package main

import (
	"fmt"
	"log"
	"net/http"

	// -----------------------------------------------------------
	// IMPORTANT: Local imports use the full module path
	// Module Name: "github.com/wes-and-me/api-server-project"
	// -----------------------------------------------------------
	// "github.com/wes-and-me/api-server-project/pkg/types"
	// "github.com/wes-and-me/api-server-project/pkg/utils"
	"github.com/wesleywinston/wds/pkg/types"
	"github.com/wesleywinston/wds/pkg/utils"

	"github.com/gorilla/mux"
)

// The root route handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Use the local types package to create an object
	sampleUser := types.NewUser("u123", "Wesley", "wesley@example.com")

	// 2. Use the local utils package to format a string
	message := utils.FormatMessage(sampleUser.Name)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\nUser ID: %s", message, sampleUser.ID)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define a simple route
	r.HandleFunc("/", homeHandler).Methods("GET")

	// Start the server
	port := ":8080"
	log.Printf("Starting server on port %s...", port)

	// This ensures your external dependency (mux) is included in go.sum
	// You should run this command in your terminal after creating the files:
	// go mod tidy

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
