package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func (s *Service) createContainer(ctx context.Context, image string, environment map[string]string, secrets map[string]string) (string, error) {
	env := []string{}

	for k, v := range environment {
		env = append(env, fmt.Sprintf("%v=%v", k, v))
	}

	for k, v := range secrets {
		env = append(env, fmt.Sprintf("%v=%v", k, v))
	}

	containerCfg := &container.Config{
		Image:     image,
		OpenStdin: true,
		StdinOnce: true,
		Env:       env,
	}

	hostConfig := &container.HostConfig{}
	nwConfig := &network.NetworkingConfig{}
	pf := &specs.Platform{}

	body, err := s.client.ContainerCreate(ctx, containerCfg, hostConfig, nwConfig, pf, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container. err: %w", err)
	}

	return body.ID, nil
}
