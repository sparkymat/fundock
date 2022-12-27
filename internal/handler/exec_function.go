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
		executionStartedAt := time.Now()

		invocationID, err := db.CreateInvocation(
			c.Request().Context(),
			*fn,
			&input,
			executionStartedAt,
		)
		if err != nil {
			c.Logger().Errorf("db.CreateInvocation failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create invocation db entry")
		}

		output, err := dockerSvc.Run(c.Request().Context(), fn.Image, input)
		if err != nil {
			c.Logger().Errorf("docker.Run failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to run function")
		}

		executionEndedAt := time.Now()

		err = db.UpdateInvocation(c.Request().Context(), *invocationID, &output, executionEndedAt.Sub(executionStartedAt).Milliseconds())
		if err != nil {
			c.Logger().Errorf("db.UpdateInvocation failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update invocation db entry")
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/fn/%v", fn.Name))
	}
}
