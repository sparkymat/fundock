package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
)

func (s *Service) pullImage(ctx context.Context, image string) error {
	resp, err := s.client.ImagePull(ctx, image, types.ImagePullOptions{
		All: true,
	})
	if err != nil {
		return fmt.Errorf("failed to pull image. err: %w", err)
	}

	_, err = io.ReadAll(resp)
	if err != nil {
		return fmt.Errorf("failed to read image pull response. err: %w", err)
	}

	return nil
}
