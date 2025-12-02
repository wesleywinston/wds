package models

import (
	"fmt"
	"time"
)

// =================================================================
// 1. Define the AccountEntity Interface
// =================================================================

// AccountEntity defines the contract for any entity (Vendor or Buyer)
// that a User can be associated with. This allows the User struct to
// hold a reference to either without knowing the concrete type.
type AccountEntity interface {
	GetLicenseID() string
	IsCompliant() bool
	String() string
}

type Vendor struct {
	ID                    string                  `json:"id"`
	BusinessName          string                  `json:"businessName"`
	OKStateLicenseID      string                  `json:"okStateLicenseID"`
	LicenseExpirationDate time.Time               `json:"licenseExpirationDate"`
	Status                AccountStatus           `json:"status"`
	ComplianceStatus      AccountComplianceStatus `json:"complianceStatus"`
	ContactInfo           ContactInfo             `json:"contactInfo"`
	MenuEnabled           bool                    `json:"menuEnabled"` // Flag to show/hide the Vendor's products on the marketplace.
	CreatedAt             time.Time               `json:"createdAt"`
	VendorCatalog		  Catalog				  `json:"catalog"`
	// ID, BusinessName, OKStateLicenseID, LicenseExpirationDate, ComplianceStatus, ContactInfo, MenuEnabled, CreatedAt
	// add a Users[] object that represents the accounts for each business, i.e Easy Street has Wes, Brighton, Chad, etc.
}

//
// {
	// {
		// "Blue Dream",
		// "item_32489348": {

		// }
	// }

	// }
// }
//

type Buyer struct {
	ID                    string                  `json:"id"`
	BusinessName          string                  `json:"businessName"`
	OKStateLicenseID      string                  `json:"okStateLicenseID"`
	LicenseExpirationDate time.Time               `json:"licenseExpirationDate"`
	Status                AccountStatus           `json:"status"`
	ComplianceStatus      AccountComplianceStatus `json:"complianceStatus"`
	ContactInfo           ContactInfo             `json:"contactInfo"`
	CreatedAt             time.Time               `json:"createdAt"`
	// Users                 []User      `json:"users"`
	// MenuEnabled           bool        `json:"menuEnabled"` // Flag to show/hide the Vendor's products on the marketplace.
}

// Vendor implements the AccountEntity interface.
func (v Vendor) GetLicenseID() string {
	return v.OKStateLicenseID
}

func (v Vendor) IsCompliant() bool {
	return v.ComplianceStatus == "VERIFIED" && time.Now().Before(v.LicenseExpirationDate)
}

func (v Vendor) String() string {
	return fmt.Sprintf("Vendor: %s (License: %s)", v.BusinessName, v.OKStateLicenseID)
}

// Buyer implements the AccountEntity interface.
func (b Buyer) GetLicenseID() string {
	return b.OKStateLicenseID
}

func (b Buyer) IsCompliant() bool {
	return b.ComplianceStatus == "VERIFIED" && time.Now().Before(b.LicenseExpirationDate)
}

func (b Buyer) String() string {
	return fmt.Sprintf("Buyer: %s (License: %s)", b.BusinessName, b.OKStateLicenseID)
}

// ContactInfo represents a contact entity in our system.
type ContactInfo struct {
	FullName []string `json:"fullName"`
	Address  string   `json:"address"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
}

type Name struct {
	FirstName string `json:"firstName"` // FullName.FirstName
	LastName  string `json:"lastName"`
}

type OrderItem struct {
	ProductID string  `json:"productID"` // Foreign key to the Product document.
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // Unit price at time of order (to account for price changes, promos. and deals)
}

type Timeline struct {
	PlacedAt    string `json:"placedAt"`
	AcceptedAt  string `json:"acceptedAt"`
	ShippedAt   string `json:"shippedAt"`
	DeliveredAt string `json:"deliveredAt"`
	CancelledAt string `json:"cancelledAt"`
	CompletedAt string `json:"completedAt"`
}

// User represents a user entity in our system.
// Fields are exported (capitalized) so they can be accessed
// and manipulated by other packages like main.
type User struct {
	ID           string   `json:"id"`           // Unique Firestore/Database ID.
	FullName     []string `json:"fullName"`     // Full name of the user.
	Email        string   `json:"email"`        // Email address.
	PasswordHash string   `json:"passwordHash"` // Hashed password.
	// FirstName          string `json:"firstName"`          // First name of the user.
	// LastName           string `json:"lastName"`           // Last name of the user.
	Role               UserRole      `json:"role"`               // Defines permissions and UI views. // coorelats with type Vendor and type Buyer
	Status             AccountStatus `json:"status"`             // Account status, linked to license validity.
	AccountData        AccountEntity `json:"accountData"`        // 'BUYER', 'VENDOR', 'ADMIN'
	AssociatedEntityID string        `json:"associatedEntityID"` // Foreign key reference to the Vendor or Buyer document this user belongs to.
}

type Order struct {
	ID                  string      `json:"id"`
	BuyerID             string      `json:"buyerID"`  // Foreign key to the User document.
	VendorID            string      `json:"vendorID"` // Foreign key to the Vendor document.
	Status              string      `json:"status"`   // originates from type OrderStatus (PENDING, ACCEPTED, PROCESSING, SHIPPED, DELIVERED, CANCELLED, COMPLETED)
	OrderStatusTimeline Timeline    `json:"orderStatusTimeline"`
	Items               []OrderItem `json:"items"`
	SubTotal            float64     `json:"subTotal"`      // sum of all line items
	ExciseTax           float64     `json:"exciseTax"`     // sum of all line items * excise tax rate
	SalesTax            float64     `json:"salesTax"`      // sum of all line items * sales tax rate
	TotalPrice          float64     `json:"totalPrice"`    // sum of all line items + excise tax + sales tax
	PaymentStatus       string      `json:"paymentStatus"` // 'PENDING', 'PAID', 'REFUNDED'
	PaymentMethod       string      `json:"paymentMethod"` // 'CASH', 'CREDIT_CARD', 'DEBIT_CARD', 'CHECK', 'OTHER'
	ShippingAddress     string      `json:"shippingAddress"`
	ShippingCost        float64     `json:"shippingCost"`
	// PlacedAt            string      `json:"placedAt"` // timestamp as a string
	// placedAt, acceptedAt, shippedAt, deliveredAt, cancelledAt, completedAt
	DeliveryDate string `json:"deliveryDate"` // Scheduled delivery date or pickup
	UpdatedAt    string `json:"updatedAt"`    // timestamp as a string
}

// --- Helper Struct for External API Response ---

// OMMAVerificationResponse simulates the data returned from an external compliance API.
type OMMAVerificationResponse struct {
	LicenseID      string    `json:"license_id"`
	IsActive       bool      `json:"is_active"`
	ExpirationDate time.Time `json:"expiration_date"`
	EntityType     string    `json:"entity_type"` // e.g., "Grower", "Dispensary", "Processor
}

type UserRole string
type AccountStatus string
type AccountComplianceStatus string // compliance status of business
type OrderStatus string
type PaymentStatus string // Tracks B2B payment status.
type PaymentMethod string // Tracks B2B payment method.
