package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// -----------------------------------------------------------
	// IMPORTANT: Local imports use the full module path
	// Module Name: "github.com/wes-and-me/api-server-project"
	// -----------------------------------------------------------
	// "github.com/wes-and-me/api-server-project/pkg/types"
	// "github.com/wes-and-me/api-server-project/pkg/utils"
	// "github.com/wesleywinston/wds/pkg/models"
	"github.com/wesleywinston/wds/pkg/models"
	"github.com/wesleywinston/wds/pkg/utils"

	"github.com/gorilla/mux"
)

// The root route handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Create a Vendor entity
	// Buyer is derived from the license type. Vendor is derived from the license type.
	// User ID and vendorID are seperately generated.
	vendorEntity := models.Vendor{
		ID:                    "vendor_213",
		BusinessName:          "Green Harvest Farms",
		OKStateLicenseID:      "PAAA-DJ3F-28JJ-283H", // G - grow , P - processor , D - dispensary
		LicenseExpirationDate: time.Now().AddDate(0, 3, 0), // Expires in 3 months
		Status:                "",
		ComplianceStatus:      "VERIFIED",
		ContactInfo: models.ContactInfo{
			FullName: []string{"Brighton", "Haney"},
			Address:  "123 Main St.",
			Email:    "brighton@friendlymarket.net",
			Phone:    "555-555-5555",
		},
		// GrowerLicenseType: "Tier 2",
	}

	// 2. Use the local types package to create an object
	// id, name, email string, passwordHash string, firstName string, lastName string, role string, status string, associatedEntityID string
	sampleUser := models.NewUser(
		"user_23473247234",
		[]string{"Brighton", "Haney"},
		"brighton@friendlymarket.net",
		"",
		// "Brighton", // take first and last name from fields in the UI, and create a name object combining the two fields. firstName is always 0, lastName is always 1
		// "Haney",SS
		"Admin",
		"Active",
		vendorEntity, // AccountEntity // account data
		"",           // AssociatedEntityID
	)

	// 2. Use the local utils package to format a string
	message := utils.FormatMessage(sampleUser.FullName)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\nUser ID: %s", message, sampleUser.ID)
}

// SignupHandler handles the creation of a new user account.
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /auth/signup")

	// Note: We no longer need to check r.Method == http.MethodPost here, 
	// as Gorilla Mux has already enforced it in main.go.

	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding signup payload: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// --- TO DO: IMPLEMENT ACCOUNT CREATION LOGIC ---
	// Placeholder Success Response:
	response := models.AuthResponse{
		Message: fmt.Sprintf("User %s successfully created as a %s. (Logic TBD)", req.Email, req.Role),
		UserID:  "mock-user-1234",
		Token:   "mock-jwt-token-...",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	log.Printf("Signup attempt successful for: %s", req.Email)
}

// LoginHandler authenticates an existing user.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /auth/login")

	// Note: We no longer need to check r.Method == http.MethodPost here.

	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding login payload: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// --- TO DO: IMPLEMENT LOGIN LOGIC ---
	// Placeholder Success Response:
	response := models.AuthResponse{
		Message: fmt.Sprintf("Welcome back, %s. Successfully logged in. (Logic TBD)", req.Email),
		Token:   "new-mock-jwt-token-...",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("Login attempt successful for: %s", req.Email)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// routes
	//
	// / (default homepage)
	// /product/:id
	// /vendor/dashboard
	// auth/login
	// auth/signup
	// /vault
	// community/feed
	//

	// internal database scrape for updating strain list
	// run once a week and pull new info

	// Define a simple route
	r.HandleFunc("/", homeHandler).Methods("GET")

	//
	// --- AUTHENTICATION ROUTES ---
	// Endpoint: POST /auth/signup
	// We use .Methods("POST") to ensure this handler only runs for POST requests.
	r.HandleFunc("/auth/signup", SignupHandler).Methods("POST")
	
	// Endpoint: POST /auth/login
	// We use .Methods("POST") to ensure this handler only runs for POST requests.
	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")

	// --- HEALTH CHECK ROUTE ---
	// This will respond to GET requests on /health
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "API is up and running!")
	}).Methods("GET")

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
