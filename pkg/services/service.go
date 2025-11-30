package services

import (
	"context"
	"errors"
	"time"

	"github.com/wesleywinston/wds/pkg/models"
)

var (
	ErrLicenseExpired  = errors.New("license is expired based on state records")
	ErrLicenseInactive = errors.New("license is inactive or invalid based on state records")
)

// VerifyLicenseExternally simulates the API call to the state (OMMA Verify).
// In a real application, this would be an HTTP call to a licensed third-party verification service
// or a government API endpoint.
func VerifyLicenseExternally(ctx context.Context, licenseID string) (*models.OMMAVerificationResponse, error) {
	// --- SIMULATION LOGIC START ---
	// Real-world implementation would involve using net/http to call a secure endpoint.

	// Simulate connection failure (e.g., network error)
	if licenseID == "OMMA-FAIL" {
		return nil, errors.New("external compliance API call failed due to timeout")
	}

	// Simulate an active license expiring 6 months from now
	if licenseID == "OMMA-ACTIVE-VENDOR" || licenseID == "OMMA-ACTIVE-BUYER" {
		response := &models.OMMAVerificationResponse{
			LicenseID:      licenseID,
			IsActive:       true,
			ExpirationDate: time.Now().AddDate(0, 6, 0), // Valid for 6 more months
			EntityType:     "Dispensary",
		}
		return response, nil
	}

	// Simulate an expired license
	if licenseID == "OMMA-EXPIRED" {
		response := &models.OMMAVerificationResponse{
			LicenseID:      licenseID,
			IsActive:       false,
			ExpirationDate: time.Now().AddDate(-1, 0, 0), // Expired 1 year ago
			EntityType:     "Grower",
		}
		return response, ErrLicenseExpired
	}

	// Simulate an invalid/non-existent license
	return nil, ErrLicenseInactive
	// --- SIMULATION LOGIC END ---
}

// CheckInternalLicenseStatus performs high-frequency checks on our own database record.
// This is used for all transactional requests (e.g., placing an order, adding a product).
func CheckInternalLicenseStatus(entityID string, licenseExpiry time.Time, status models.Status) error {
	// 1. Check our cached compliance status
	if status != "VERIFIED" {
		return errors.New("internal status check failed: entity is not marked as VERIFIED")
	}

	// 2. Check the expiration date locally (fast check)
	if time.Now().After(licenseExpiry) {
		return errors.New("internal status check failed: license has expired")
	}

	return nil // License is internally valid and active
}
