package presenter

import "github.com/sparkymat/fundock/model"

//nolint:tagliatelle
type Invocation struct {
	ID          string  `json:"id"`
	StartedTime *string `json:"started_time"`
	EndedTime   *string `json:"ended_time"`
	Output      *string `json:"output"`
}

//nolint:tagliatelle
type InvocationsList struct {
	PageNumber uint32       `json:"page_number"`
	PageSize   uint32       `json:"page_size"`
	Items      []Invocation `json:"items"`
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

func InvocationsListFromModels(pageNumber uint32, pageSize uint32, invocations []model.Invocation) InvocationsList {
	presentedList := InvocationsList{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Items:      []Invocation{},
	}

	for _, inv := range invocations {
		presentedList.Items = append(presentedList.Items, InvocationFromModel(inv))
	}

	return presentedList
}
