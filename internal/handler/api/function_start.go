package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/auth"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
	"github.com/sparkymat/fundock/services/runner"
)

func FunctionStart(cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientName, _ := c.Get(auth.ClientNameKey).(string)

		name := c.Param("name")

		var requestBody bytes.Buffer

		defer c.Request().Body.Close()

		_, err := io.Copy(&requestBody, c.Request().Body)
		if err != nil {
			//nolint:wrapcheck
			return c.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error": "failed to read request body",
			})
		}

		functionRunner, err := runner.New(cfg, db, dockerSvc)
		if err != nil {
			c.Logger().Warnf("failed to create runner. err: %w", err)

			//nolint:wrapcheck
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to initialize runner",
			})
		}

		fn, invocationID, err := functionRunner.SetupFunction(c.Request().Context(), name, clientName, requestBody.String())
		if err != nil {
			//nolint:wrapcheck
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":         "failed to setup function",
				"internalError": err.Error(),
			})
		}

		//nolint:errcheck
		go functionRunner.ExecFunction(c.Request().Context(), fn, *invocationID, requestBody.String())

		invocation, err := db.FetchInvocation(c.Request().Context(), *invocationID)
		if err != nil {
			//nolint:wrapcheck
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": "failed to run function",
			})
		}

		presentedInvocation := presenter.InvocationFromModel(*invocation)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedInvocation)
	}
}
