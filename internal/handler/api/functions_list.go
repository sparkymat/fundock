package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
)

type FunctionsListInput struct {
	PageSize   uint32 `query:"page_size"`
	PageNumber uint32 `query:"page_number"`
	Query      string `query:"query"`
}

func FunctionsList(_ configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &FunctionsListInput{}
		if err := c.Bind(input); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "invalid input")
		}

		functions, err := db.FetchFunctions(c.Request().Context(), input.PageNumber, input.PageSize)
		if err != nil {
			c.Logger().Errorf("db.FetchFunctions failed with err: %v", err)

			return renderError(c, http.StatusInternalServerError, "failed to load functions")
		}

		presentedFunctionsList := presenter.FunctionsListFromModels(input.PageNumber, input.PageSize, functions)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedFunctionsList)
	}
}
