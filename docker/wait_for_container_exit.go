package docker

import (
	"context"
	"time"
)

const ContainerTimeoutSeconds = 5

func (s *Service) waitForContainerExit(ctx context.Context, containerID string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ContainerTimeoutSeconds*time.Second)
	defer cancel()

	resultC, errC := s.client.ContainerWait(ctx, containerID, "")
	select {
	case err := <-errC:
		return 0, err
	case result := <-resultC:
		return result.StatusCode, nil
	}
}
