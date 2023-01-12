package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
	"github.com/sparkymat/fundock/model"
)

type InvocationsListInput struct {
	PageSize   uint32 `query:"page_size"`
	PageNumber uint32 `query:"page_number"`
	Function   string `query:"fn"`
}

func InvocationsList(_ configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &InvocationsListInput{}
		if err := c.Bind(input); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "invalid input")
		}

		var invocations []model.Invocation

		if input.Function != "" {
			fn, err := db.FetchFunction(c.Request().Context(), input.Function)
			if err != nil {
				c.Logger().Errorf("db.FetchFunction failed with err: %v", err)

				return renderError(c, http.StatusInternalServerError, "failed to load function")
			}

			invocations, err = db.FetchFunctionInvocations(c.Request().Context(), fn.ID, input.PageNumber, input.PageSize)
			if err != nil {
				c.Logger().Errorf("db.FetchFunctionInvocations failed with err: %v", err)

				return renderError(c, http.StatusInternalServerError, "failed to load function invocations")
			}
		} else {
			var err error

			invocations, err = db.FetchInvocations(c.Request().Context(), input.PageNumber, input.PageSize)
			if err != nil {
				c.Logger().Errorf("db.FetchInvocations failed with err: %v", err)

				return renderError(c, http.StatusInternalServerError, "failed to load invocations")
			}
		}

		presentedInvocationsList := presenter.InvocationsListFromModels(input.PageNumber, input.PageSize, invocations)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedInvocationsList)
	}
}
