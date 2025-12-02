package models

// AuthRequest is a structure to decode incoming JSON payloads for signup/login
type AuthRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role,omitempty"` 
}

// AuthResponse is a standard structure for returning a success message or token
type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	UserID  string `json:"userId,omitempty"`
}