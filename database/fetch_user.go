package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchUser(ctx context.Context, username string) (*model.User, error) {
	sqlString := `SELECT
	u.*
FROM users u
WHERE u.username = $1
`

	var user model.User

	if err := s.conn.QueryRowxContext(ctx, sqlString, username).StructScan(&user); err != nil {
		return nil, fmt.Errorf("failed to query for user. err: %w", err)
	}

	return &user, nil
}
