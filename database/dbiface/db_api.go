package dbiface

import (
	"context"

	"github.com/sparkymat/fundock/model"
)

type DBAPI interface {
	AutoMigrate() error
	FetchFunctions(ctx context.Context, query string, pageSize uint32, pageNumber uint32) ([]model.Function, error)
}
