package runner

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
)

var (
	ErrDBFailure = errors.New("db failure")
)

func New(cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) (*Service, error) {
	return &Service{
		cfg:       cfg,
		db:        db,
		dockerSvc: dockerSvc,
	}, nil
}

type Service struct {
	cfg       configiface.ConfigAPI
	db        dbiface.DBAPI
	dockerSvc dockeriface.DockerAPI
}

func (s *Service) ExecFunction(ctx context.Context, functionName string, input string) (*string, *string, error) {
	// Fetch the function
	fn, err := s.db.FetchFunction(ctx, functionName)
	if err != nil || fn == nil {
		if err == nil {
			err = ErrDBFailure
		}

		return nil, nil, fmt.Errorf("function fetch failed. err: %w", err)
	}

	var loggedInput *string

	if !fn.SkipLogging {
		loggedInput = &input
	}

	// Start recording the invocation
	invocationID, err := s.db.CreateInvocation(
		ctx,
		*fn,
		loggedInput,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("create invocation record failed. err: %w", err)
	}

	executionStartedAt := time.Now()

	// Update invocation starting
	err = s.db.UpdateInvocationStarted(ctx, *invocationID, executionStartedAt)
	if err != nil {
		return nil, nil, fmt.Errorf("update 'invocation started' failed. err: %w", err)
	}

	// Run image with input
	output, err := s.dockerSvc.Run(ctx, fn.Image, input)
	executionEndedAt := time.Now()

	if err != nil {
		errorMessage := err.Error()

		// Update invocation failed
		err = s.db.UpdateInvocationFailed(ctx, *invocationID, executionEndedAt, &errorMessage)
		if err != nil {
			return nil, nil, fmt.Errorf("update 'invocation failed' failed. err: %w", err)
		}

		return invocationID, &output, nil
	}

	var loggedOutput *string

	if !fn.SkipLogging {
		copiedOutput := output
		loggedOutput = &copiedOutput
	}

	// Update invocation succeeded
	err = s.db.UpdateInvocationSucceeded(ctx, *invocationID, executionEndedAt, loggedOutput)
	if err != nil {
		return nil, nil, fmt.Errorf("update 'invocation succeeded' failed. err: %w", err)
	}

	return invocationID, &output, nil
}
