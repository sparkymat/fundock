package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/model"
	"github.com/sparkymat/fundock/presenter"
	"github.com/sparkymat/fundock/view"
)

func InvocationShow(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusNotFound, "Invocation not found")
		}

		inv, err := db.FetchInvocation(c.Request().Context(), id)
		if err != nil {
			c.Logger().Errorf("db.FetchInvocation failed with err: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load invocation")
		}

		presentedInv := presenter.InvocationFromModel(*inv)

		var fn *model.Function
		var presentedFn *presenter.Function

		if inv.FunctionID.Valid {
			fn, err = db.FetchFunction(c.Request().Context(), inv.FunctionName)
			if err != nil {
				fn = nil
				c.Logger().Errorf("db.FetchInvocation failed with err: %v", err)
			} else {
				pf := presenter.FunctionFromModel(*fn)
				presentedFn = &pf
			}
		}

		pageHTML := view.InvocationShow(presentedInv, presentedFn)
		htmlString := view.Layout(fmt.Sprintf("fundock | %v | %v", inv.ID, inv.FunctionName), pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
