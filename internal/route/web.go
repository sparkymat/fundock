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
	app := e.Group("")

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	app.GET("/", handler.Home(cfg, db))
	app.GET("/functions", handler.Functions(cfg, db))
	app.GET("/fn/:name", handler.FunctionShow(cfg, db))
	app.POST("/exec/:name", handler.ExecFunction(cfg, db, dockerSvc))
	app.GET("/invocations/:id", handler.InvocationShow(cfg, db))
	app.GET("/functions/new", handler.NewFunction(cfg, db))
	app.POST("/functions", handler.CreateFuncion(cfg, db))
	app.GET("/api_tokens", handler.APITokens(cfg, db))
	app.POST("/api_tokens", handler.CreateAPIToken(cfg, db))
}
