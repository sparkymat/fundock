package model

import (
	"database/sql"
	"time"
)

type InvocationStatus string

const (
	InvocationStatusPending   = "pending"
	InvocationStatusRunning   = "running"
	InvocationStatusFailed    = "failed"
	InvocationStatusSucceeded = "succeeded"
)

type Invocation struct {
	ID           string           `db:"id"`
	Status       InvocationStatus `db:"status"`
	FunctionName string           `db:"function_name"`
	FunctionID   sql.NullString   `db:"function_id"`
	Image        string           `db:"image"`
	Input        sql.NullString   `db:"input"`
	Output       sql.NullString   `db:"output"`
	ErrorMessage sql.NullString   `db:"error_message"`
	CreatedAt    time.Time        `db:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at"`
	StartedAt    *time.Time       `db:"started_at"`
	EndedAt      *time.Time       `db:"ended_at"`
}
