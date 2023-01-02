package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID                string         `db:"id"`
	Username          string         `db:"username"`
	EncryptedPassword string         `db:"encrypted_password"`
	Name              sql.NullString `db:"name"`
	Email             sql.NullString `db:"email"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at"`
}
