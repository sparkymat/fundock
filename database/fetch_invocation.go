package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) FetchInvocation(ctx context.Context, id string) (*model.Invocation, error) {
	sqlString := `SELECT
	inv.*
FROM invocations inv
WHERE inv.id = $1
`

	var inv model.Invocation

	if err := s.conn.QueryRowxContext(ctx, sqlString, id).StructScan(&inv); err != nil {
		return nil, fmt.Errorf("failed to query for invocation. err: %w", err)
	}

	return &inv, nil
}
