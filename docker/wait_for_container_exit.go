package docker

import (
	"context"
	"time"
)

func (s *Service) waitForContainerExit(ctx context.Context, containerID string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resultC, errC := s.client.ContainerWait(ctx, containerID, "")
	select {
	case err := <-errC:
		return 0, err
	case result := <-resultC:
		return result.StatusCode, nil
	}
}
