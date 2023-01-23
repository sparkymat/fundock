package runner

import (
	"context"
	"fmt"
	"time"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) ExecFunction(ctx context.Context, fn *model.Function, invocationID string, input string, environment map[string]string, secrets map[string]string) (*string, error) {
	executionStartedAt := time.Now()

	// Update invocation starting
	err := s.db.UpdateInvocationStarted(ctx, invocationID, executionStartedAt)
	if err != nil {
		return nil, fmt.Errorf("update 'invocation started' failed. err: %w", err)
	}

	// Run image with input
	output, err := s.dockerSvc.Run(ctx, fn.Image, input, environment, secrets)
	executionEndedAt := time.Now()

	if err != nil {
		errorMessage := err.Error()

		// Update invocation failed
		err = s.db.UpdateInvocationFailed(ctx, invocationID, executionEndedAt, &errorMessage)
		if err != nil {
			return nil, fmt.Errorf("update 'invocation failed' failed. err: %w", err)
		}

		return &output, nil
	}

	var loggedOutput *string

	if !fn.SkipLogging {
		copiedOutput := output
		loggedOutput = &copiedOutput
	}

	// Update invocation succeeded
	err = s.db.UpdateInvocationSucceeded(ctx, invocationID, executionEndedAt, loggedOutput)
	if err != nil {
		return nil, fmt.Errorf("update 'invocation succeeded' failed. err: %w", err)
	}

	return &output, nil
}
