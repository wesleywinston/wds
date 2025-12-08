package handlers

import (
	"github.com/wesleywinston/wds/pkg/models"
)

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