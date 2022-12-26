package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
)

func (s *Service) getContainerLogs(ctx context.Context, containerID string) (string, error) {
	reader, err := s.client.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", fmt.Errorf("failed to get container log stream. err: %w", err)
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("failed to read logs. err: %w", err)
	}

	return string(body), nil
}
