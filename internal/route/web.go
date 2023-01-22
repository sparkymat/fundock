package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/fundock/auth"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler"
)

func setupWebRoutes(e *echo.Echo, cfg configiface.ConfigAPI, db dbiface.DBAPI) {
	webApp := e.Group("")

	webApp.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "login") //nolint:wrapcheck
	})

	authenticatedWebApp := webApp.Group("")

	if cfg.SingleUser() {
		authenticatedWebApp.Use(auth.AdminTokenInjector(cfg, db))
	}

	authenticatedWebApp.Use(auth.TokenOrSessionAuthMiddleware(cfg, db))

	authenticatedWebApp.GET("/", handler.Home(cfg, db))
}
