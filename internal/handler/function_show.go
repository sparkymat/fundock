package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/presenter"
	"github.com/sparkymat/fundock/view"
)

func FunctionShow(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := GetCSRFToken(c)

		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusNotFound, "Function not found")
		}

		fn, err := db.FetchFunction(c.Request().Context(), name)
		if err != nil || fn == nil {
			c.Logger().Errorf("db.FetchFunction failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load function")
		}

		latestInvocations, err := db.FetchFunctionInvocations(c.Request().Context(), fn.ID, 1, 10)
		if err != nil || fn == nil {
			c.Logger().Errorf("db.FetchFunctionInvocations failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load function invocations")
		}

		presentedFn := presenter.FunctionFromModel(*fn)
		presentedInvocations := []presenter.Invocation{}
		for _, in := range latestInvocations {
			presentedInvocations = append(presentedInvocations, presenter.InvocationFromModel(in))
		}

		pageHTML := view.FunctionShow(csrfToken, presentedFn, presentedInvocations)
		htmlString := view.Layout(fmt.Sprintf("fundock | %v", fn.Name), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
