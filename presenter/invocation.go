package presenter

import (
	"fmt"

	"github.com/sparkymat/fundock/model"
)

type Invocation struct {
	Timestamp string
	Duration  string
}

func InvocationFromModel(in model.Invocation) Invocation {
	return Invocation{
		Timestamp: in.ExecutedAt.String(),
		Duration:  fmt.Sprintf("%d ms", in.ExecDurationMilliSeconds.Int64),
	}
}
