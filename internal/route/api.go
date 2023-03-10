package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/fundock/auth"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
	"github.com/sparkymat/fundock/internal/handler/api"
)

func setupAPIRoutes(e *echo.Echo, cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) {
	apiApp := e.Group("/api")

	apiApp.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	apiApp.Use(auth.TokenOrSessionAuthMiddleware(cfg, db))

	apiApp.POST("/functions", api.FunctionCreate(cfg, db))
	apiApp.GET("/functions", api.FunctionsList(cfg, db))
	apiApp.GET("/fn/:name", api.FunctionShow(cfg, db))

	apiApp.POST("/fn/:name/exec", api.FunctionExec(cfg, db, dockerSvc))
	apiApp.POST("/fn/:name/start", api.FunctionStart(cfg, db, dockerSvc))

	apiApp.GET("/invocations", api.InvocationsList(cfg, db))
	apiApp.GET("/invocations/:id", api.InvocationShow(cfg, db))

	apiApp.GET("/api_tokens", api.APITokensList(cfg, db))
	apiApp.POST("/api_tokens", api.APITokenCreate(cfg, db))
}
