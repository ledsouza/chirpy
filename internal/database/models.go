// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	UserID    uuid.UUID
}

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdateAt  time.Time
	UserID    uuid.UUID
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

type User struct {
	ID             uuid.UUID
	UpdatedAt      time.Time
	CreatedAt      time.Time
	Email          string
	HashedPassword string
}
