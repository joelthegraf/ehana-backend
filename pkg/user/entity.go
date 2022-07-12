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
