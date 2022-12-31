package presenter

import (
	"fmt"

	"github.com/sparkymat/fundock/model"
)

type Invocation struct {
	ID           string
	FunctionName string
	Image        string
	Status       string
	Timestamp    string
	Duration     string
	ClientName   string
	Input        *string
	Output       *string
}

func InvocationFromModel(in model.Invocation) Invocation {
	presentedIn := Invocation{
		ID:           in.ID,
		FunctionName: in.FunctionName,
		Image:        in.Image,
		Status:       string(in.Status),
		ClientName:   in.ClientName,
	}

	if in.StartedAt != nil {
		presentedIn.Timestamp = in.StartedAt.String()

		if in.EndedAt != nil {
			presentedIn.Duration = fmt.Sprintf("%d ms", in.EndedAt.Sub(*in.StartedAt).Milliseconds())
		}
	}

	if in.Input.Valid {
		presentedIn.Input = &in.Input.String
	}

	if in.Output.Valid {
		presentedIn.Output = &in.Output.String
	}

	return presentedIn
}
