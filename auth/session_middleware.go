package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/jwt"
)

const UserKey = "user"

const (
	sessionName = "fundock_session"
	tokenKey    = "auth_token"
)

var (
	ErrTokenMissing = errors.New("token missing")
)

func SessionMiddleware(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(cfg, db, c)
			if err != nil {
				return handleAuthFailure(c)
			}

			return next(c)
		}
	}
}

func sessionAuth(cfg configiface.ConfigAPI, db dbiface.DBAPI, c echo.Context) error {
	authToken := getSessionValue(c, tokenKey)
	if authToken == "" {
		return ErrTokenMissing
	}

	username, err := jwt.ValidateTokenString(cfg.JWTSecret(), authToken)
	if err != nil {
		return err
	}

	user, err := db.FetchUser(c.Request().Context(), *username)
	if err != nil {
		return err
	}

	c.Set(UserKey, user)

	return nil
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
