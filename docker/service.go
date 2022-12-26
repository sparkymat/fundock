package docker

import (
	"context"
	"fmt"
	"io"

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

	err = s.createContainer(ctx, image, input)

	return "", err
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

func (s *Service) createContainer(ctx context.Context, image string, input string) error {
	containerCfg := &container.Config{
		Image: image}
	hostConfig := &container.HostConfig{
		AutoRemove: true,
	}
	nwConfig := &network.NetworkingConfig{}
	pf := &specs.Platform{}

	body, err := s.client.ContainerCreate(ctx, containerCfg, hostConfig, nwConfig, pf, "")
	if err != nil {
		return fmt.Errorf("failed to create container. err: %w", err)
	}

	fmt.Printf("id=%v\n", body.ID)

	return nil
}
