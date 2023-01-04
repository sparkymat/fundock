package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func SessionMiddleware(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cfg.SingleUser() {
				user, err := db.FetchUser(c.Request().Context(), "admin")
				if err != nil {
					c.Logger().Warnf("failed to load admin user in single user mode. err: %v", err)

					return echo.NewHTTPError(http.StatusUnauthorized, "Failed to load admin user")
				}

				r := c.Request()
				updatedContext := context.WithValue(c.Request().Context(), ContextUserKey, user)
				updatedRequest := r.WithContext(updatedContext)
				c.SetRequest(updatedRequest)

				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
