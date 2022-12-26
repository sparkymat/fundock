package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
)

func (s *Service) getContainerLogs(ctx context.Context, containerID string) (string, error) {
	reader, err := s.client.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", err
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
