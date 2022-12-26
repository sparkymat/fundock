package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func (s *Service) createContainer(ctx context.Context, image string, input string) (string, error) {
	containerCfg := &container.Config{
		Image: image}
	hostConfig := &container.HostConfig{}
	nwConfig := &network.NetworkingConfig{}
	pf := &specs.Platform{}

	body, err := s.client.ContainerCreate(ctx, containerCfg, hostConfig, nwConfig, pf, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container. err: %w", err)
	}

	return body.ID, nil
}
