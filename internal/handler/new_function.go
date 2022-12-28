package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/view"
)

func NewFunction(_ configiface.ConfigAPI, _ dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := GetCSRFToken(c)

		pageHTML := view.NewFunction(csrfToken)
		htmlString := view.Layout("fundock | new", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
