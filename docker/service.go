package docker

import (
	"context"
	"fmt"

	dockerlib "github.com/docker/docker/client"
)

func New() (*Service, error) {
	client, err := dockerlib.NewClientWithOpts(dockerlib.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize docker client. err: %w", err)
	}

	return &Service{client: client}, nil
}

type Service struct {
	client *dockerlib.Client
}

func (s *Service) Run(ctx context.Context, image string, input string) (string, error) {

	err := s.pullImage(ctx, image)
	if err != nil {
		return "", err
	}

	containerID, err := s.createContainer(ctx, image, input)
	if err != nil {
		return "", err
	}

	err = s.runContainer(ctx, containerID)
	if err != nil {
		return "", err
	}

	_, err = s.waitForContainerExit(ctx, containerID)
	if err != nil {
		return "", err
	}

	response, err := s.getContainerLogs(ctx, containerID)
	if err != nil {
		return "", err
	}

	return response, err
}
