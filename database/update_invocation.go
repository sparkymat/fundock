package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) UpdateInvocationStarted(ctx context.Context, id string, startedAt time.Time) error {
	sqlString := `UPDATE invocations
	SET
		status = $1,
		started_at = $2
	WHERE id = $3;
`
	if _, err := s.conn.ExecContext(ctx, sqlString, string(model.InvocationStatusRunning), startedAt, id); err != nil {
		return fmt.Errorf("failed to update invocation started. err: %w", err)
	}

	return nil
}

func (s *Service) UpdateInvocationSucceeded(ctx context.Context, id string, endedAt time.Time, output *string) error {
	sqlString := `UPDATE invocations
	SET
		status = $1,
		output = $2,
		ended_at = $3
	WHERE id = $4;
`
	if _, err := s.conn.ExecContext(ctx, sqlString, string(model.InvocationStatusSucceeded), output, endedAt, id); err != nil {
		return fmt.Errorf("failed to update invocation succeeded. err: %w", err)
	}

	return nil
}

func (s *Service) UpdateInvocationFailed(ctx context.Context, id string, endedAt time.Time, errorMessage *string) error {
	sqlString := `UPDATE invocations
	SET
		status = $1,
		error_message = $2,
		ended_at = $3
	WHERE id = $4;
`
	if _, err := s.conn.ExecContext(ctx, sqlString, string(model.InvocationStatusFailed), errorMessage, endedAt, id); err != nil {
		return fmt.Errorf("failed to update invocation failed. err: %w", err)
	}

	return nil
}
