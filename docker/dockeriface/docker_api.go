package dockeriface

import "context"

type DockerAPI interface {
	Run(ctx context.Context, image string, input string) (string, error)
}
