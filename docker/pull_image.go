package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
)

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
