package presenter

import (
	"fmt"

	"github.com/sparkymat/fundock/model"
	"github.com/xeonx/timeago"
)

type Function struct {
	URL              string
	Name             string
	Image            string
	SkipLogging      bool
	CreatedTimestamp string
}

func FunctionFromModel(fn model.Function) Function {
	return Function{
		Name:             fn.Name,
		Image:            fn.Image,
		URL:              fmt.Sprintf("/fn/%s", fn.Name),
		SkipLogging:      fn.SkipLogging,
		CreatedTimestamp: timeago.English.Format(fn.CreatedAt),
	}
}
