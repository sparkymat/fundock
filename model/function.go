package model

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx/types"
)

type Function struct {
	ID          string         `db:"id"`
	Name        string         `db:"name"`
	Image       string         `db:"image"`
	SkipLogging bool           `db:"skip_logging"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	Environment types.JSONText `db:"environment"`
	Secrets     types.JSONText `db:"secrets"`
}

func (f Function) EnvironmentJSON() (map[string]string, error) {
	jsonMap := map[string]string{}

	err := f.Environment.Unmarshal(&jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal environment. err: %w", err)
	}

	return jsonMap, nil
}

func (f Function) SecretsJSON() (map[string]string, error) {
	jsonMap := map[string]string{}

	err := f.Secrets.Unmarshal(&jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secrets. err: %w", err)
	}

	return jsonMap, nil
}
