package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/fundock/auth"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
	"github.com/sparkymat/fundock/internal/handler"
)

func setupWebRoutes(e *echo.Echo, cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) {
	webApp := e.Group("")

	webApp.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "login")
	})

	authenticatedWebApp := webApp.Group("")
	authenticatedWebApp.Use(auth.SessionMiddleware(cfg, db))

	authenticatedWebApp.GET("/", handler.Home(cfg, db))
	authenticatedWebApp.GET("/functions", handler.Functions(cfg, db))
	authenticatedWebApp.GET("/fn/:name", handler.FunctionShow(cfg, db))
	authenticatedWebApp.POST("/exec/:name", handler.ExecFunction(cfg, db, dockerSvc))
	authenticatedWebApp.GET("/invocations/:id", handler.InvocationShow(cfg, db))
	authenticatedWebApp.GET("/functions/new", handler.NewFunction(cfg, db))
	authenticatedWebApp.POST("/functions", handler.CreateFuncion(cfg, db))
	authenticatedWebApp.GET("/api_tokens", handler.APITokens(cfg, db))
	authenticatedWebApp.POST("/api_tokens", handler.CreateAPIToken(cfg, db))
}
