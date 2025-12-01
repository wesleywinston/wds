package models

// NewUser is a simple constructor function (exported).
func NewUser(id string, fullName []string, email string, passwordHash string, role string, status string, accountData AccountEntity, associatedEntityID string) User {
	return User{
		ID:           id,
		FullName:     fullName,
		Email:        email,
		PasswordHash: passwordHash,
		// FirstName:          firstName,
		// LastName:           lastName,
		Role:               UserRole(role),
		Status:             AccountStatus(status),
		AccountData:        accountData,
		AssociatedEntityID: associatedEntityID,
	}
}
