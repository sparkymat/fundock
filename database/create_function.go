package database

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Service) CreateFunction(
	ctx context.Context,
	name string,
	image string,
	skipLogging bool,
	environment map[string]string,
	secrets map[string]string,
) (*string, error) {
	sqlString := `INSERT INTO functions
	(name, image, skip_logging, environment, secrets)
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING id
	`

	environmentString, err := json.Marshal(environment)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal environment. err: %w", err)
	}

	secretsString, err := json.Marshal(secrets)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal environment. err: %w", err)
	}

	var functionID string
	if err := s.conn.QueryRowContext( //nolint:execinquery
		ctx,
		sqlString,
		name,
		image,
		skipLogging,
		environmentString,
		secretsString,
	).Scan(
		&functionID,
	); err != nil {
		return nil, fmt.Errorf("failed to create function. err: %w", err)
	}

	return &functionID, nil
}
