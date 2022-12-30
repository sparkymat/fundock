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
	apiApp.Use(auth.TokenAuthMiddleware(cfg, db))

	apiApp.POST("/exec/:name", api.ExecFunction(cfg, db, dockerSvc))
}
