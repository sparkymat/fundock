package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/view"
)

func Home(_ configiface.ConfigAPI, _ dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		pageHTML := view.Home()
		htmlString := view.Layout("fundock", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
