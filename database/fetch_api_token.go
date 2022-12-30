package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchAPIToken(ctx context.Context, tokenString string) (*model.APIToken, error) {
	sqlString := `SELECT
	t.*
FROM api_tokens t
WHERE t.token = $1
`

	var apiToken model.APIToken

	if err := s.conn.QueryRowxContext(ctx, sqlString, tokenString).StructScan(&apiToken); err != nil {
		return nil, fmt.Errorf("failed to query for api token. err: %w", err)
	}

	return &apiToken, nil
}
