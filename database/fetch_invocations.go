package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchInvocations(ctx context.Context, pageNumber uint32, pageSize uint32) ([]model.Invocation, error) {
	if pageSize == 0 || pageNumber == 0 {
		return nil, ErrInvalidPagination
	}

	offset := (pageNumber - 1) * pageSize

	sqlString := `SELECT
	i.*
FROM invocations i
ORDER BY COALESCE(i.ended_at, i.started_at, i.created_at) DESC
OFFSET $1
LIMIT $2
`

	invocations := []model.Invocation{}

	rows, err := s.conn.QueryxContext(ctx, sqlString, offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to run db query. err: %w", err)
	}

	for rows.Next() {
		var inv model.Invocation

		err := rows.StructScan(&inv)
		if err != nil {
			return nil, fmt.Errorf("failed to scan db result. err: %w", err)
		}

		invocations = append(invocations, inv)
	}

	return invocations, nil
}
