package database

import (
	"context"
	"fmt"
)

func (s *Service) CreateUser(ctx context.Context, username string, encryptedPassword string, email *string, name *string) (*string, error) {
	sqlString := `INSERT INTO users
	(username, encrypted_password, email, name)
	VALUES
	($1, $2, $3, $4)
	RETURNING id
	`

	var userID string
	if err := s.conn.QueryRowContext( //nolint:execinquery
		ctx,
		sqlString,
		username,
		encryptedPassword,
		email,
		name,
	).Scan(
		&userID,
	); err != nil {
		return nil, fmt.Errorf("failed to create user. err: %w", err)
	}

	return &userID, nil
}
