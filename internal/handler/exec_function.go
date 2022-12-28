package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
)

func ExecFunction(cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusNotFound, "Function not found")
		}

		fn, err := db.FetchFunction(c.Request().Context(), name)
		if err != nil || fn == nil {
			c.Logger().Errorf("db.FetchFunction failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load function")
		}

		var inputBuffer bytes.Buffer

		_, err = io.Copy(&inputBuffer, c.Request().Body)
		if err != nil {
			c.Logger().Errorf("db.FetchFunction failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load function")
		}

		input := inputBuffer.String()

		invocationID, err := db.CreateInvocation(
			c.Request().Context(),
			*fn,
			&input,
		)
		if err != nil {
			c.Logger().Errorf("db.CreateInvocation failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create invocation db entry")
		}

		executionStartedAt := time.Now()

		err = db.UpdateInvocationStarted(c.Request().Context(), *invocationID, executionStartedAt)
		if err != nil {
			c.Logger().Errorf("db.UpdateInvocationStarted failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation started")
		}

		output, err := dockerSvc.Run(c.Request().Context(), fn.Image, input)
		executionEndedAt := time.Now()
		if err != nil {
			errorMessage := err.Error()
			err = db.UpdateInvocationFailed(c.Request().Context(), *invocationID, executionEndedAt, &errorMessage)
			if err != nil {
				c.Logger().Errorf("db.UpdateInvocationFailed failed with err: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation failed")
			}

			c.Logger().Errorf("docker.Run failed with err: %v", err)
			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", fn.Name))
		}

		err = db.UpdateInvocationSucceeded(c.Request().Context(), *invocationID, executionEndedAt, &output)
		if err != nil {
			c.Logger().Errorf("db.UpdateInvocationSucceeded failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation succeeded")
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", fn.Name))
	}
}
