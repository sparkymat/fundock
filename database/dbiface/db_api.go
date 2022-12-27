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
	CreateInvocation(ctx context.Context, fn model.Function, input *string, executedAt time.Time) (*string, error)
	UpdateInvocation(ctx context.Context, id string, output *string, executionTimeMS int64) error
}
