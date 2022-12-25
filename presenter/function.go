package presenter

import (
	"fmt"

	"github.com/sparkymat/fundock/model"
)

type Function struct {
	URL   string
	Name  string
	Image string
}

func FunctionFromModel(fn model.Function) Function {
	return Function{
		Name:  fn.Name,
		Image: fn.Image,
		URL:   fmt.Sprintf("/fn/%s", fn.Name),
	}
}
