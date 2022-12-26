package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
)

func (s *Service) attachContainer(ctx context.Context, containerID string, input string) error {
	resp, err := s.client.ContainerAttach(ctx, containerID, types.ContainerAttachOptions{
		Stdin:  true,
		Stream: true,
	})
	if err != nil {
		return fmt.Errorf("failed to attach to running container. err: %w", err)
	}
	defer resp.Close()

	_, err = resp.Conn.Write([]byte(input))
	if err != nil {
		return fmt.Errorf("failed to write to container stdin. err: %w", err)
	}

	return nil
}
