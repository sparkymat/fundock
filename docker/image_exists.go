package docker

import (
	"context"
	"fmt"
	"regexp"

	"github.com/docker/docker/api/types"
)

func (s *Service) imageExists(ctx context.Context, image string) (bool, error) {
	strippedName := stripHostFromImage(image)

	summaries, err := s.client.ImageList(ctx, types.ImageListOptions{All: false})
	if err != nil {
		return false, fmt.Errorf("failed to load image list. err: %w", err)
	}

	for _, imageSummary := range summaries {
		img, _, err := s.client.ImageInspectWithRaw(ctx, imageSummary.ID)
		if err != nil {
			continue
		}
		for _, repoTag := range img.RepoTags {
			if repoTag == strippedName {
				return true, nil
			}
		}
	}

	return false, nil
}

func stripHostFromImage(image string) string {
	hostPrefixRegex := regexp.MustCompile(`^\w+(\.\w+)+\/`)
	if hostPrefixRegex.MatchString(image) {
		return hostPrefixRegex.ReplaceAllString(image, "")
	}

	return image
}
