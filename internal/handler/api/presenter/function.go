package presenter

import "github.com/sparkymat/fundock/model"

//nolint:tagliatelle
type Function struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	SkipLogging bool              `json:"skip_logging"`
	CreatedTime string            `json:"created_time"`
	Environment map[string]string `json:"environment"`
	Secrets     []string          `json:"secrets"`
}

//nolint:tagliatelle
type FunctionsList struct {
	PageNumber uint32     `json:"page_number"`
	PageSize   uint32     `json:"page_size"`
	Items      []Function `json:"items"`
}

func FunctionFromModel(fn model.Function) Function {
	presentedFn := Function{
		ID:          fn.ID,
		Name:        fn.Name,
		Image:       fn.Image,
		SkipLogging: fn.SkipLogging,
		CreatedTime: fn.CreatedAt.String(),
		Secrets:     []string{},
	}

	environment, err := fn.EnvironmentJSON()
	if err == nil {
		presentedFn.Environment = environment
	}

	secrets, err := fn.SecretsJSON()
	if err == nil {
		for secretKey := range secrets {
			presentedFn.Secrets = append(presentedFn.Secrets, secretKey)
		}
	}

	return presentedFn
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
