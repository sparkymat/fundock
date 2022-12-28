package docker

import (
	"bytes"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
)

func (s *Service) getContainerLogs(ctx context.Context, containerID string) (string, error) {
	multiReader, err := s.client.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", fmt.Errorf("failed to get container log stream. err: %w", err)
	}
	defer multiReader.Close()

	var stdoutBuffer bytes.Buffer

	var stderrBuffer bytes.Buffer

	_, err = stdcopy.StdCopy(&stdoutBuffer, &stderrBuffer, multiReader)
	if err != nil {
		return "", fmt.Errorf("failed to read container log stream. err: %w", err)
	}

	body := stdoutBuffer.String()

	return body, nil
}
