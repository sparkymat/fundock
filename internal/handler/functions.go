package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/presenter"
	"github.com/sparkymat/fundock/view"
)

const (
	DefaultPageSize   = 20
	DefaultPageNumber = 1
)

func Functions(_ configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		functions, err := db.FetchFunctions(c.Request().Context(), DefaultPageNumber, DefaultPageSize)
		if err != nil {
			c.Logger().Errorf("db.FetchFunctions failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load functions")
		}

		presentedFunctions := []presenter.Function{}

		for _, fn := range functions {
			presentedFunctions = append(presentedFunctions, presenter.FunctionFromModel(fn))
		}

		pageHTML := view.Functions(presentedFunctions)
		htmlString := view.Layout("fundock | functions", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
