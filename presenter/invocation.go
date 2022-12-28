package presenter

import (
	"fmt"

	"github.com/sparkymat/fundock/model"
)

type Invocation struct {
	ID        string
	Status    string
	Timestamp string
	Duration  string
}

func InvocationFromModel(in model.Invocation) Invocation {
	presentedIn := Invocation{
		ID:     in.ID,
		Status: string(in.Status),
	}

	if in.StartedAt != nil {
		presentedIn.Timestamp = in.StartedAt.String()

		if in.EndedAt != nil {
			presentedIn.Duration = fmt.Sprintf("%d ms", in.EndedAt.Sub(*in.StartedAt).Milliseconds())
		}
	}

	return presentedIn
}
