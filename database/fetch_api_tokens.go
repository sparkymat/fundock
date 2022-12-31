package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchAPITokens(ctx context.Context, pageNumber uint32, pageSize uint32) ([]model.APIToken, error) {
	if pageSize == 0 || pageNumber == 0 {
		return nil, ErrInvalidPagination
	}

	offset := (pageNumber - 1) * pageSize

	sqlString := `SELECT
	t.*
FROM api_tokens t
ORDER BY t.updated_at DESC
OFFSET $1
LIMIT $2
`
	apiTokens := []model.APIToken{}

	rows, err := s.conn.QueryxContext(ctx, sqlString, offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to run db query. err: %w", err)
	}

	for rows.Next() {
		var apiToken model.APIToken

		err := rows.StructScan(&apiToken)
		if err != nil {
			return nil, fmt.Errorf("failed to scan db result. err: %w", err)
		}

		apiTokens = append(apiTokens, apiToken)
	}

	return apiTokens, nil
}
