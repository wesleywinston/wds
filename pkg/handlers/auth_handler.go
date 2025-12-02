package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wesleywinston/wds/pkg/models"
)

// NewUserRequest is the payload for creating a new platform user.
type NewUserRequest struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Role               string `json:"role"`               // e.g., models.RoleVendor
	AssociatedEntityID string `json:"associatedEntityId"` // Vendor/Buyer ID
}

// CheckEntityCompliance ensures the referenced Vendor/Buyer entity is valid and active.
// In a real app, this would query the DB for the entity and check its ComplianceStatus.
func CheckEntityCompliance(ctx context.Context, db *firestore.Client, entityID string, role string) error {
	if entityID == "" && (role == "VENDOR" || role == "BUYER") {
		return errors.New("missing associated entity ID for licensed role")
	}

	// --- SIMULATION LOGIC: Entity check against DB ---
	// Real-world: Fetch the Vendor/Buyer document and check its ComplianceStatus and ExpiryDate.
	// For now, we only allow specific IDs to simulate compliance.
	if role == models.RoleVendor && entityID != "V-ACTIVE-101" {
		return errors.New("vendor entity not found or compliance failed")
	}
	if role == models.RoleBuyer && entityID != "B-ACTIVE-202" {
		return errors.New("buyer entity not found or compliance failed")
	}
	// --- END SIMULATION ---

	return nil
}

// CreateUser handles the registration of a new user and assigns role-based permissions.
func CreateUser(db *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req NewUserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// 1. Role-based Entity and Compliance Check
		// TODO: Implement actual entity and compliance checks
		// For now, we only allow specific IDs to simulate compliance.
		//
		//	ADMIN, BUYER, VENDOR
		if req.Role == "VENDOR" || req.Role == "BUYER" {
			if err := CheckEntityCompliance(ctx, db, req.AssociatedEntityID, req.Role); err != nil {
				log.Printf("Entity compliance failed for user %s: %v", req.Email, err)
				http.Error(w, "Cannot create user: Associated business entity is invalid or non-compliant.", http.StatusForbidden)
				return
			}
		} else if req.Role == "ADMIN" {
			// Admin accounts require human intervention/approval (simulated here by checking a secret key or internal list)
			if req.Email != "admin@company.com" {
				// Prevent self-registering as admin
				http.Error(w, "Admin accounts must be provisioned internally.", http.StatusForbidden)
				return
			}
		} else {
			http.Error(w, "Invalid user role specified.", http.StatusBadRequest)
			return
		}

		// 2. Create the User Model
		newUser := models.User{
			// ID, FullName, Email, PasswordHash, Role, Status, AssociatedEntityID
			ID:                 time.Now().Format("20060102-150405-user"),
			Email:              req.Email,
			PasswordHash:       "hashed_" + req.Password, // Hashing logic goes here
			FirstName:          req.FirstName,
			LastName:           req.LastName,
			Role:               req.Role,
			AssociatedEntityID: req.AssociatedEntityID,
			Status:             models.StatusActive, // Assuming compliance check passed, otherwise PENDING_APPROVAL
			CreatedAt:          time.Now(),
		}

		// 3. Persist to Database (Simulated)
		// db.Collection("users").Doc(newUser.ID).Set(ctx, newUser)
		log.Printf("SUCCESS: User created: %s (%s) for Entity: %s", newUser.Email, newUser.Role, newUser.AssociatedEntityID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "User account created successfully.", "userID": newUser.ID})
	}
}
