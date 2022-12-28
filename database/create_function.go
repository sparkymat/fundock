package database

import (
	"context"
	"fmt"
)

func (s *Service) CreateFunction(ctx context.Context, name string, image string, skipLogging bool) (*string, error) {
	sqlString := `INSERT INTO functions
	(name, image, skip_logging)
	VALUES
	($1, $2, $3)
	RETURNING id
	`

	var functionID string
	if err := s.conn.QueryRowContext(
		ctx,
		sqlString,
		name,
		image,
		skipLogging,
	).Scan(
		&functionID,
	); err != nil {
		return nil, fmt.Errorf("failed to create function. err: %w", err)
	}

	return &functionID, nil
}
