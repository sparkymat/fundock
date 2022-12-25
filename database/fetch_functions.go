package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

var (
	ErrInvalidPagination = errors.New("invalid page request")
)

func (s *Service) FetchFunctions(ctx context.Context, query string, pageSize uint32, pageNumber uint32) ([]model.Function, error) {
	if pageSize == 0 || pageNumber == 0 {
		return nil, ErrInvalidPagination
	}

	offset := (pageNumber - 1) * pageSize

	sqlString := `SELECT
	f.id,
	f.name,
	f.image,
	f.skip_logging,
	f.created_at,
	f.updated_at
FROM functions f
WHERE f.name ILIKE $1
OFFSET $2
LIMIT $3
ORDER BY name desc
`
	functions := []model.Function{}

	rows, err := s.conn.QueryxContext(ctx, sqlString, fmt.Sprintf("%%%s%%", query), offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to run db query. err: %w", err)
	}

	for rows.Next() {
		var fn model.Function
		err := rows.StructScan(&fn)
		if err != nil {
			return nil, fmt.Errorf("failed to scan db result. err: %w", err)
		}
		functions = append(functions, fn)
	}

	return functions, nil
}
