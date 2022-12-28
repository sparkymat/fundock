package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
)

//nolint:funlen,revive,cyclop
func ExecFunction(_ configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusNotFound, "Function not found")
		}

		input := c.FormValue("input")

		// Fetch the function
		fn, err := db.FetchFunction(c.Request().Context(), name)
		if err != nil || fn == nil {
			c.Logger().Errorf("db.FetchFunction failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load function")
		}

		var loggedInput *string

		if !fn.SkipLogging {
			copiedInput := input
			loggedInput = &copiedInput
		}

		// Start recording the invocation
		invocationID, err := db.CreateInvocation(
			c.Request().Context(),
			*fn,
			loggedInput,
		)
		if err != nil {
			c.Logger().Errorf("db.CreateInvocation failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create invocation db entry")
		}

		executionStartedAt := time.Now()

		// Update invocation starting
		err = db.UpdateInvocationStarted(c.Request().Context(), *invocationID, executionStartedAt)
		if err != nil {
			c.Logger().Errorf("db.UpdateInvocationStarted failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation started")
		}

		// Run image with input
		output, err := dockerSvc.Run(c.Request().Context(), fn.Image, input)
		executionEndedAt := time.Now()

		if err != nil {
			errorMessage := err.Error()

			// Update invocation failed
			err = db.UpdateInvocationFailed(c.Request().Context(), *invocationID, executionEndedAt, &errorMessage)
			if err != nil {
				c.Logger().Errorf("db.UpdateInvocationFailed failed with err: %v", err)

				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation failed")
			}

			c.Logger().Errorf("docker.Run failed with err: %v", err)

			//nolint:wrapcheck
			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", fn.Name))
		}

		var loggedOutput *string

		if !fn.SkipLogging {
			copiedOutput := output
			loggedOutput = &copiedOutput
		}

		// Update invocation succeeded
		err = db.UpdateInvocationSucceeded(c.Request().Context(), *invocationID, executionEndedAt, loggedOutput)
		if err != nil {
			c.Logger().Errorf("db.UpdateInvocationSucceeded failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation succeeded")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", fn.Name))
	}
}
