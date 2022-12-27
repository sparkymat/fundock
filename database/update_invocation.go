package database

import (
	"context"
	"fmt"
)

func (s *Service) UpdateInvocation(ctx context.Context, id string, output *string, executionTimeMS int64) error {
	sqlString := `UPDATE invocations
	SET
		output = $1,
		execution_time_ms = $2
	WHERE id = $3;
`
	if _, err := s.conn.ExecContext(ctx, sqlString, output, executionTimeMS, id); err != nil {
		return fmt.Errorf("failed to update invocation. err: %w", err)
	}

	return nil
}
