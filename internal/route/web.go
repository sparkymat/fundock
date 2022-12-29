package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	webApp.GET("/", handler.Home(cfg, db))
	webApp.GET("/functions", handler.Functions(cfg, db))
	webApp.GET("/fn/:name", handler.FunctionShow(cfg, db))
	webApp.POST("/exec/:name", handler.ExecFunction(cfg, db, dockerSvc))
	webApp.GET("/invocations/:id", handler.InvocationShow(cfg, db))
	webApp.GET("/functions/new", handler.NewFunction(cfg, db))
	webApp.POST("/functions", handler.CreateFuncion(cfg, db))
	webApp.GET("/api_tokens", handler.APITokens(cfg, db))
	webApp.POST("/api_tokens", handler.CreateAPIToken(cfg, db))
}
