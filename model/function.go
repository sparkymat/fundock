package model

import "time"

type Function struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Image       string    `db:"image"`
	SkipLogging bool      `db:"skip_logging"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
