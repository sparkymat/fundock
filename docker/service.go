package docker

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	dockerlib "github.com/docker/docker/client"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
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

func (s *Service) getContainerLogs(ctx context.Context, containerID string) (string, error) {
	reader, err := s.client.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{})
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

func (s *Service) pullImage(ctx context.Context, image string) error {
	resp, err := s.client.ImagePull(ctx, image, types.ImagePullOptions{
		All: true,
	})
	if err != nil {
		return err
	}

	_, err = io.ReadAll(resp)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) createContainer(ctx context.Context, image string, input string) (string, error) {
	containerCfg := &container.Config{
		Image: image}
	hostConfig := &container.HostConfig{
		AutoRemove: true,
	}
	nwConfig := &network.NetworkingConfig{}
	pf := &specs.Platform{}

	body, err := s.client.ContainerCreate(ctx, containerCfg, hostConfig, nwConfig, pf, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container. err: %w", err)
	}

	return body.ID, nil
}

func (s *Service) runContainer(ctx context.Context, containerID string) error {
	err := s.client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return fmt.Errorf("failed to create container. err: %w", err)
	}

	return nil
}
