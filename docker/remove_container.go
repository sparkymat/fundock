package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
)

func (s *Service) removeContainer(ctx context.Context, containerID string) error {
	err := s.client.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return fmt.Errorf("failed to remove container. err: %w", err)
	}

	return nil
}
