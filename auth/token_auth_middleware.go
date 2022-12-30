package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
)

const ClientNameKey = "client_name"

func TokenAuthMiddleware(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("X-API-Key")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "token not found",
				})
			}

			apiToken, err := db.FetchAPIToken(c.Request().Context(), tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "token load failed",
				})
			}

			c.Set(ClientNameKey, apiToken.ClientName)

			return next(c)
		}
	}
}
