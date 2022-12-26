package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
)

func ExecFunction(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusNotFound, "Function not found")
		}

		fn, err := db.FetchFunction(c.Request().Context(), name)
		if err != nil || fn == nil {
			c.Logger().Errorf("db.FetchFunction failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load function")
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", fn.Name))
	}
}
