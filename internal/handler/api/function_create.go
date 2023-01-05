package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
)

type FunctionCreateInput struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	SkipLogging bool   `json:"skip_logging"`
}

func FunctionCreate(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &FunctionCreateInput{}
		if err := c.Bind(input); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "invalid input")
		}

		_, err := db.CreateFunction(c.Request().Context(), input.Name, input.Image, input.SkipLogging)
		if err != nil {
			c.Logger().Errorf("db.CreateFunction failed with err: %v", err)

			return renderError(c, http.StatusInternalServerError, "function creation failed")
		}

		fn, err := db.FetchFunction(c.Request().Context(), input.Name)
		if err != nil {
			c.Logger().Errorf("db.FetchFunction failed with err: %v", err)

			return renderError(c, http.StatusInternalServerError, "function fetch failed")
		}

		presentedFn := presenter.FunctionFromModel(*fn)

		return c.JSON(http.StatusOK, presentedFn)
	}
}
