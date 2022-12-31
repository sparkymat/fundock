package database

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) CreateInvocation(ctx context.Context, fn model.Function, clientName string, input *string) (*string, error) {
	sqlString := `INSERT INTO invocations
	(function_name, function_id, image, client_name, input)
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING id
	`

	var invocationID string
	if err := s.conn.QueryRowContext( //nolint:execinquery
		ctx,
		sqlString,
		fn.Name,
		fn.ID,
		fn.Image,
		clientName,
		input,
	).Scan(
		&invocationID,
	); err != nil {
		return nil, fmt.Errorf("failed to create invocation. err: %w", err)
	}

	return &invocationID, nil
}
