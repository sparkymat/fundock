package presenter

import (
	"fmt"

	"github.com/sparkymat/fundock/model"
)

type Invocation struct {
	ID        string
	Timestamp string
	Duration  string
}

func InvocationFromModel(in model.Invocation) Invocation {
	return Invocation{
		ID:        in.ID,
		Timestamp: in.ExecutedAt.String(),
		Duration:  fmt.Sprintf("%d ms", in.ExecDurationMilliSeconds.Int64),
	}
}
