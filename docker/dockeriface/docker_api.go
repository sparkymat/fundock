package dockeriface

import "context"

type DockerAPI interface {
	Run(ctx context.Context, image string, input string, environment map[string]string, secrets map[string]string) (string, error)
}
