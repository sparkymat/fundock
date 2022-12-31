package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/presenter"
)

func InvocationShow(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			//nolint:wrapcheck
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "record not found",
			})
		}

		invocation, err := db.FetchInvocation(c.Request().Context(), id)
		if err != nil {
			c.Logger().Errorf("db.FetchInvocation failed with err: %v", err)

			//nolint:wrapcheck
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "record not found",
			})
		}

		presentedInvocation := presenter.InvocationFromModel(*invocation)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedInvocation)
	}
}
