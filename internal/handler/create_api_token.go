package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
)

func CreateAPIToken(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientName := c.FormValue("client_name")
		token := strings.ToLower(ulid.Make().String())

		_, err := db.CreateAPIToken(c.Request().Context(), clientName, token)
		if err != nil {
			c.Logger().Errorf("db.CreateAPIToken failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create api token")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/api_tokens")
	}
}
