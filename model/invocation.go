package model

import (
	"database/sql"
	"time"
)

type Invocation struct {
	ID                       string         `db:"id"`
	FunctionName             string         `db:"function_name"`
	FunctionID               sql.NullString `db:"function_id"`
	Image                    string         `db:"image"`
	Input                    sql.NullString `db:"input"`
	Output                   sql.NullString `db:"output"`
	CreatedAt                time.Time      `db:"created_at"`
	UpdatedAt                time.Time      `db:"updated_at"`
	ExecutedAt               *time.Time     `db:"executed_at"`
	ExecDurationMilliSeconds sql.NullInt64  `db:"exec_duration_ms"`
}
