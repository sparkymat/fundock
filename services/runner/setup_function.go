package runner

import (
	"context"
	"fmt"

	"github.com/sparkymat/fundock/model"
)

func (s *Service) SetupFunction(ctx context.Context, functionName string, clientName string, input string) (*model.Function, *string, error) {
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
		clientName,
		loggedInput,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("create invocation record failed. err: %w", err)
	}

	return fn, invocationID, nil
}
