package models

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
