package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
	"github.com/sparkymat/fundock/services/runner"
)

func ExecFunction(cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusNotFound, "Function not found")
		}

		input := c.FormValue("input")

		functionRunner, err := runner.New(cfg, db, dockerSvc)
		if err != nil {
			c.Logger().Warnf("failed to create runner. err: %w", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to intiialize runner")
		}

		fn, invocationID, err := functionRunner.SetupFunction(c.Request().Context(), name, "web", input)
		if err != nil {
			c.Logger().Warnf("failed to setup function. err: %w", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to setup function")
		}

		_, err = functionRunner.ExecFunction(c.Request().Context(), fn, *invocationID, input)
		if err != nil {
			c.Logger().Warnf("failed to exec function. err: %w", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to exec function")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/invocations/%s", *invocationID))
	}
}
