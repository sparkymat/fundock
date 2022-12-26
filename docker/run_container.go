package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
)

func (s *Service) runContainer(ctx context.Context, containerID string) error {
	err := s.client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return fmt.Errorf("failed to create container. err: %w", err)
	}

	return nil
}
