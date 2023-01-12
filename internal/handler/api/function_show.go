package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
)

func FunctionShow(_ configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")

		fn, err := db.FetchFunction(c.Request().Context(), name)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return renderError(c, http.StatusNotFound, "function not found")
			}

			return renderError(c, http.StatusInternalServerError, "function load failed")
		}

		presentedFn := presenter.FunctionFromModel(*fn)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedFn)
	}
}
