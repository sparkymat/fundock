package auth

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LoadSession(c echo.Context) (*sessions.Session, error) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize session. err: %w", err)
	}

	return sess, nil
}

func SaveSession(c echo.Context, sess *sessions.Session, tokenString *string) error {
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 8, //nolint:gomnd
		HttpOnly: true,
	}

	if tokenString != nil {
		sess.Values[tokenKey] = *tokenString
	} else {
		sess.Values[tokenKey] = ""
	}

	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("failed to save session. err: %w", err)
	}

	return nil
}
