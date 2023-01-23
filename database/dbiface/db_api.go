package dbiface

import (
	"context"
	"time"

	"github.com/sparkymat/fundock/model"
)

//nolint:interfacebloat
type DBAPI interface {
	AutoMigrate() error
	FetchFunctions(ctx context.Context, pageSize uint32, pageNumber uint32) ([]model.Function, error)
	FetchFunction(ctx context.Context, name string) (*model.Function, error)
	CreateInvocation(ctx context.Context, fn model.Function, clientName string, input *string) (*string, error)
	UpdateInvocationStarted(ctx context.Context, id string, startedAt time.Time) error
	UpdateInvocationSucceeded(ctx context.Context, id string, endedAt time.Time, output *string) error
	UpdateInvocationFailed(ctx context.Context, id string, endedAt time.Time, errorMessage *string) error
	FetchFunctionInvocations(ctx context.Context, functionID string, pageNumber uint32, pageSize uint32) ([]model.Invocation, error)
	FetchInvocation(ctx context.Context, id string) (*model.Invocation, error)
	CreateFunction(ctx context.Context, name string, image string, skipLogging bool, environment map[string]string, secrets map[string]string) (*string, error)
	FetchAPITokens(ctx context.Context, pageNumber uint32, pageSize uint32) ([]model.APIToken, error)
	CreateAPIToken(ctx context.Context, clientName string, token string) (*string, error)
	DeleteAPIToken(ctx context.Context, id string) error
	FetchAPIToken(ctx context.Context, tokenString string) (*model.APIToken, error)
	FetchAPITokenByID(ctx context.Context, tokenString string) (*model.APIToken, error)
	FetchUser(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, username string, encryptedPassword string, email *string, name *string) (*string, error)
	FetchInvocations(ctx context.Context, pageNumber uint32, pageSize uint32) ([]model.Invocation, error)
}
