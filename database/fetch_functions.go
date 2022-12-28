package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

var ErrInvalidPagination = errors.New("invalid page request")

func (s *Service) FetchFunctions(ctx context.Context, pageNumber uint32, pageSize uint32) ([]model.Function, error) {
	if pageSize == 0 || pageNumber == 0 {
		return nil, ErrInvalidPagination
	}

	offset := (pageNumber - 1) * pageSize

	sqlString := `SELECT
	f.*
FROM functions f
ORDER BY f.name DESC
OFFSET $1
LIMIT $2
`
	functions := []model.Function{}

	rows, err := s.conn.QueryxContext(ctx, sqlString, offset, pageSize)
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
