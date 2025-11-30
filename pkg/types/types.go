package types

// User represents a user entity in our system.
// Fields are exported (capitalized) so they can be accessed
// and manipulated by other packages like main.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewUser is a simple constructor function (exported).
func NewUser(id, name, email string) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
