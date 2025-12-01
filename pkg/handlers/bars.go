package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore" // Placeholder for Firestore client
	"github.com/wesleywinston/wds/pkg/models"
	"github.com/wesleywinston/wds/pkg/services"
)

// VendorRegistrationRequest is the payload expected from the frontend.
type VendorRegistrationRequest struct {
	BusinessName     string `json:"businessName"`
	OkStateLicenseID string `json:"okStateLicenseId"`
	// ... other required vendor fields
}

// RegisterVendor handles the initial registration and license verification for a new Vendor.
func RegisterVendor(db *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req VendorRegistrationRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// --- STEP 1: External Real-Time License Verification ---
		// We verify the license ID against the state's (simulated) API.
		ommaResponse, err := services.VerifyLicenseExternally(ctx, req.OkStateLicenseID)
		if err != nil {
			// This covers network failure or the license being inactive/expired externally.
			log.Printf("External verification failed for %s: %v", req.OkStateLicenseID, err)
			http.Error(w, "License verification failed or license is inactive/expired.", http.StatusForbidden)
			return
		}

		// Ensure the external license check returned a response that is active.
		if !ommaResponse.IsActive {
			http.Error(w, "License is not currently active according to state records.", http.StatusForbidden)
			return
		}

		// --- STEP 2: Create Vendor Entity ---
		newVendor := models.Vendor{
			// ID, BusinessName, OKStateLicenseID, LicenseExpirationDate, ComplianceStatus, ContactInfo, MenuEnabled, CreatedAt
			ID:                    time.Now().Format("20060102-150405"), // Simplified ID generation
			BusinessName:          req.BusinessName,
			OKStateLicenseID:      req.OkStateLicenseID,
			LicenseExpirationDate: ommaResponse.ExpirationDate, // Use the date returned from the state API
			ComplianceStatus:      "VERIFIED",                  // Mark VERIFIED because the external check succeeded
			ContactInfo:           models.ContactInfo{ /* populate contact */ },
			MenuEnabled:           false,
			CreatedAt:             time.Now(),
		}

		// --- STEP 3: Persist to Database ---
		// In a real scenario, you'd use db.Collection("vendors").Doc(newVendor.ID).Set(ctx, newVendor)
		// For simulation, we'll just log success.
		log.Printf("SUCCESS: Vendor %s registered with license active until %s", newVendor.BusinessName, newVendor.LicenseExpirationDate.Format("2006-01-02"))

		// Respond to the user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Vendor registration successful. Please proceed to create your user account."})
	}
}

// CheckEntityActive is a general middleware/handler used for transactional checks (e.g., before placing an order).
func CheckEntityActive(db *firestore.Client, entityID string) error {
	ctx := context.Background()

	// --- STEP 1: Retrieve Vendor/Buyer data from our DB ---
	// For simplicity, we assume we fetch a Vendor here, but this logic would be dynamic based on role.
	var entity models.User // check direct reference of call to this function from user input

	// Simulated database retrieval: Fetch the Vendor based on the entityID
	// In reality: docSnap, err := db.Collection("vendors").Doc(entityID).Get(ctx)
	// For simulation, we assume retrieval succeeded with known good data:
	if entityID == "V-12345" {
		entity.AccountData.ComplianceStatus = "VERIFIED"
		// Set expiration 1 year from now
		entity.AccountData.LicenseExpirationDate = time.Now().AddDate(1, 0, 0)
	} else {
		// Simulate expired/bad data
		entity.AccountData.ComplianceStatus = "EXPIRED"
		entity.AccountData.LicenseExpirationDate = time.Now().AddDate(-1, 0, 0)
	}

	// --- STEP 2: Internal Compliance Check ---
	err := services.CheckInternalLicenseStatus(entityID, entity.AccountData.LicenseExpirationDate, entity.AccountData.ComplianceStatus) // this might need to be updated to entity.AccountData.(Vendor/Buyer)
	// if buy, ok := buyerUser.AccountData.(Buyer); ok {
	// 	if buy.ComplianceStatus != "VERIFIED" {
	// 		return errors.New("buyer account is not verified")
	// 	}
	// }
	if err != nil {
		// Log the failure for administrative review
		log.Printf("Transactional license check failed for entity %s: %v", entityID, err)

		// If the internal check fails, we might trigger a background job to re-verify externally
		// to catch recently renewed licenses that we haven't updated yet.
		// For now, we return the internal failure.
		return errors.New("entity license is not active or has expired")
	}

	return nil // Transactional check passed
}
