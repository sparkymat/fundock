package database

import (
	"context"
	"fmt"
)

func (s *Service) DeleteAPIToken(ctx context.Context, id string) error {
	sqlString := `DELETE FROM api_tokens
	WHERE id = $1;
`
	if _, err := s.conn.ExecContext(ctx, sqlString, id); err != nil {
		return fmt.Errorf("failed to delete api token. err: %w", err)
	}

	return nil
}
