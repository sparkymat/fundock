package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) CreateInvocation(ctx context.Context, fn model.Function, input *string, executedAt time.Time) (*string, error) {
	sqlString := `INSERT INTO invocations
	(function_name, function_id, image, input, executed_at)
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING id
	`

	var invocationID string
	if err := s.conn.QueryRowContext(
		ctx,
		sqlString,
		fn.Name,
		fn.ID,
		fn.Image,
		input,
		executedAt,
	).Scan(
		&invocationID,
	); err != nil {
		return nil, fmt.Errorf("failed to create invocation. err: %w", err)
	}

	return &invocationID, nil
}
