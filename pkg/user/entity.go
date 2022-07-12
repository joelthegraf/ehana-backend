package user

import "github.com/google/uuid"

// A User represents a ehana account.
type User struct {
	ID       uint64
	UUID     uuid.UUID
	Username string
	Password string
	Created  int64
	Modified int64
}
