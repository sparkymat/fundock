package dbiface

import (
	"context"
	"time"

	"github.com/sparkymat/fundock/model"
)

type DBAPI interface {
	AutoMigrate() error
	FetchFunctions(ctx context.Context, pageSize uint32, pageNumber uint32) ([]model.Function, error)
	FetchFunction(ctx context.Context, name string) (*model.Function, error)
	CreateInvocation(ctx context.Context, fn model.Function, input *string) (*string, error)
	UpdateInvocationStarted(ctx context.Context, id string, startedAt time.Time) error
	UpdateInvocationSucceeded(ctx context.Context, id string, endedAt time.Time, output *string) error
	UpdateInvocationFailed(ctx context.Context, id string, endedAt time.Time, errorMessage *string) error
	FetchFunctionInvocations(ctx context.Context, functionID string, pageNumber uint32, pageSize uint32) ([]model.Invocation, error)
	FetchInvocation(ctx context.Context, id string) (*model.Invocation, error)
	CreateFunction(ctx context.Context, name string, image string, skipLogging bool) (*string, error)
	FetchAPITokens(ctx context.Context, pageNumber uint32, pageSize uint32) ([]model.APIToken, error)
}
