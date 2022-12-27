package database

import (
	"context"
	"fmt"
)

func (s *Service) UpdateInvocation(ctx context.Context, id string, output *string, execDurationMS int64) error {
	sqlString := `UPDATE invocations
	SET
		output = $1,
		exec_duration_ms = $2
	WHERE id = $3;
`
	if _, err := s.conn.ExecContext(ctx, sqlString, output, execDurationMS, id); err != nil {
		return fmt.Errorf("failed to update invocation. err: %w", err)
	}

	return nil
}
