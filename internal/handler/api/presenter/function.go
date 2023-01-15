package presenter

import "github.com/sparkymat/fundock/model"

//nolint:tagliatelle
type Function struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	CreatedTime string `json:"created_time"`
}

//nolint:tagliatelle
type FunctionsList struct {
	PageNumber uint32     `json:"page_number"`
	PageSize   uint32     `json:"page_size"`
	Items      []Function `json:"items"`
}

func FunctionFromModel(fn model.Function) Function {
	return Function{
		ID:          fn.ID,
		Name:        fn.Name,
		Image:       fn.Name,
		CreatedTime: fn.CreatedAt.String(),
	}
}

func FunctionsListFromModels(pageNumber uint32, pageSize uint32, functions []model.Function) FunctionsList {
	presentedList := FunctionsList{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Items:      []Function{},
	}

	for _, fn := range functions {
		presentedList.Items = append(presentedList.Items, FunctionFromModel(fn))
	}

	return presentedList
}
