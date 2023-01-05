package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/jwt"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

const (
	sessionName = "fundock_session"
	tokenKey    = "auth_token"
)

func SessionMiddleware(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authToken := getSessionValue(c, tokenKey)
			if authToken == "" {
				return handleAuthFailure(c)
			}

			username, err := jwt.ValidateTokenString(cfg.JWTSecret(), authToken)
			if err != nil {
				return handleAuthFailure(c)
			}

			user, err := db.FetchUser(c.Request().Context(), *username)
			if err != nil {
				return handleAuthFailure(c)
			}

			r := c.Request()
			updatedContext := context.WithValue(c.Request().Context(), ContextUserKey, user)
			updatedRequest := r.WithContext(updatedContext)
			c.SetRequest(updatedRequest)

			return next(c)
		}
	}
}

func getSessionValue(c echo.Context, key string) string {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return ""
	}

	stringInterface, isString := sess.Values[key]

	if !isString {
		return ""
	}

	return stringInterface.(string) //nolint:forcetypeassert
}

func handleAuthFailure(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/login") //nolint:wrapcheck
}
