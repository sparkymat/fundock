package model

import "time"

type APIToken struct {
	ID         string     `db:"id"`
	Token      string     `db:"token"`
	ClientName string     `db:"client_name"`
	LastUsedAt *time.Time `db:"last_used_at"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}
