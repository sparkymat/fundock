package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchFunction(ctx context.Context, name string) (*model.Function, error) {
	sqlString := `SELECT
	f.*
FROM functions f
WHERE f.name = $1
`

	var fn model.Function

	if err := s.conn.QueryRowxContext(ctx, sqlString, name).StructScan(&fn); err != nil {
		return nil, fmt.Errorf("failed to query for function. err: %w", err)
	}

	return &fn, nil
}
