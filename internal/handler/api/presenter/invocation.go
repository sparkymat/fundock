package presenter

import "github.com/sparkymat/fundock/model"

type Invocation struct {
	ID          string
	StartedTime *string
	EndedTime   *string
	Output      *string
}

func InvocationFromModel(inv model.Invocation) Invocation {
	presentedInv := Invocation{
		ID: inv.ID,
	}

	if inv.StartedAt != nil {
		startedTime := inv.StartedAt.String()
		presentedInv.StartedTime = &startedTime
	}

	if inv.EndedAt != nil {
		endedTime := inv.EndedAt.String()
		presentedInv.EndedTime = &endedTime
	}

	return presentedInv
}
