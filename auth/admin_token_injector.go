package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/jwt"
)

func AdminTokenInjector(cfg configiface.ConfigAPI, _ dbiface.DBAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := LoadSession(c)
			if err != nil {
				return next(c)
			}

			tokenString, err := jwt.GenerateToken(cfg.JWTSecret(), "admin")
			if err != nil {
				return next(c)
			}

			_ = SaveSession(c, sess, &tokenString)

			return next(c)
		}
	}
}
