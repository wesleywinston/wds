package types

// ContactInfo represents a contact entity in our system.
type ContactInfo struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type Name struct {
	FirstName string `json:"firstName"` // FullName.FirstName
	LastName  string `json:"lastName"`
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
	Role               string `json:"role"`               // Defines permissions and UI views.
	Status             string `json:"status"`             // Account status, linked to license validity.
	AssociatedEntityID string `json:"associatedEntityID"` // Foreign key reference to the Vendor or Buyer document this user belongs to.
}

type Vendor struct {
	ID                    string      `json:"id"`
	BusinessName          string      `json:"businessName"`
	OKStateLicenseID      string      `json:"okStateLicenseID"`
	LicenseExpirationDate string      `json:"licenseExpirationDate"`
	Status                string      `json:"status"`
	ContactInfo           ContactInfo `json:"contactInfo"`
	MenuEnabled           bool        `json:"menuEnabled"` // Flag to show/hide the Vendor's products on the marketplace.
	// add a Users[] object that represents the accounts for each business, i.e Easy Street has Wes, Brighton, Chad, etc.
}

type Role string
type Status string

const (
	RoleAdmin  Role = "ADMIN"
	RoleBuyer  Role = "BUYER"
	RoleVendor Role = "VENDOR"
)

const (
	StatusPending  Status = "PENDING_APPROVAL"
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
)

// NewUser is a simple constructor function (exported).
func NewUser(id string, fullName []string, email string, passwordHash string, role string, status string, associatedEntityID string) User {
	return User{
		ID:           id,
		FullName:     fullName,
		Email:        email,
		PasswordHash: passwordHash,
		// FirstName:          firstName,
		// LastName:           lastName,
		Role:               role,
		Status:             status,
		AssociatedEntityID: associatedEntityID,
	}
}
