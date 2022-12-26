package database

import (
	"context"

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
		return nil, err
	}

	return &fn, nil
}
