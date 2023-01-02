package model

import "time"

type User struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
