package docker

import (
	"context"
	"fmt"

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
	containerCfg := &container.Config{}
	hostConfig := &container.HostConfig{}
	nwConfig := &network.NetworkingConfig{}
	pf := &specs.Platform{}

	body, err := s.client.ContainerCreate(ctx, containerCfg, hostConfig, nwConfig, pf, image)
	if err != nil {
		return "", fmt.Errorf("failed to create container. err: %w", err)
	}

	fmt.Printf("id=%v\n", body.ID)

	return "", nil
}
