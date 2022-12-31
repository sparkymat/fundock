package database

import (
	"context"
	"fmt"
)

func (s *Service) CreateAPIToken(ctx context.Context, clientName string, token string) (*string, error) {
	sqlString := `INSERT INTO api_tokens
	(client_name, token)
	VALUES
	($1, $2)
	RETURNING id
	`

	var apiTokenID string
	if err := s.conn.QueryRowContext( //nolint:execinquery
		ctx,
		sqlString,
		clientName,
		token,
	).Scan(
		&apiTokenID,
	); err != nil {
		return nil, fmt.Errorf("failed to create api token. err: %w", err)
	}

	return &apiTokenID, nil
}
