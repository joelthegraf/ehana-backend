package user

import "github.com/google/uuid"

// A User represents a ehana account.
type User struct {
	ID       uint64    // ID has to be unique.
	UUID     uuid.UUID // UUID has to be unique.
	Username string    // Username has to be unique.
	Password string
	Created  int64
	Modified int64
}

// The Repository interface abstracts the by insight required functionality for database interactions with the User data structure.
// This allows too simply change between different data sources.
type Repository interface {
	// Create must insert a User into the database and update the User.ID field.
	Create(user *User) error
	// FindByUUID must find a User in the database based on the User.UUID field.
	FindByUUID(UUID uuid.UUID) (User, error)
	// FindByUsername must find a User in the database based on the User.Username field.
	FindByUsername(username string) (User, error)
	// Update must update the User.Username, User.Password, and User.Modified field of a User in the database.
	Update(user User) error
	// Delete must delete a User in the database based on User.ID.
	Delete(ID uint64)
}
