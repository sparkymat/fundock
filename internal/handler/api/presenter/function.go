package presenter

import "github.com/sparkymat/fundock/model"

type Function struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	CreatedTime string `json:"created_time"`
}

func FunctionFromModel(fn model.Function) Function {
	return Function{
		ID:          fn.ID,
		Name:        fn.Name,
		Image:       fn.Name,
		CreatedTime: fn.CreatedAt.String(),
	}
}
