package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
)

func CreateFuncion(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		image := c.FormValue("image")
		skipLoggingVal := c.FormValue("skip_logging")

		skipLogging := false
		if skipLoggingVal == "on" {
			skipLogging = true
		}

		_, err := db.CreateFunction(c.Request().Context(), name, image, skipLogging)
		if err != nil {

			c.Logger().Errorf("db.CreateFunction failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create function")
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", name))
	}
}
