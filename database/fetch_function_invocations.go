package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchFunctionInvocations(ctx context.Context, functionID string, pageNumber uint32, pageSize uint32) ([]model.Invocation, error) {
	if pageSize == 0 || pageNumber == 0 {
		return nil, ErrInvalidPagination
	}

	offset := (pageNumber - 1) * pageSize

	sqlString := `SELECT
	inv.*
FROM invocations inv
WHERE inv.function_id = $1
ORDER BY inv.updated_at DESC
OFFSET $2
LIMIT $3
`

	invocations := []model.Invocation{}

	rows, err := s.conn.QueryxContext(ctx, sqlString, functionID, offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to run db query. err: %w", err)
	}

	for rows.Next() {
		var in model.Invocation

		err := rows.StructScan(&in)
		if err != nil {
			return nil, fmt.Errorf("failed to scan db result. err: %w", err)
		}

		invocations = append(invocations, in)
	}

	return invocations, nil
}
