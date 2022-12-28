package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/presenter"
	"github.com/sparkymat/fundock/view"
)

func APITokens(cfg configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiTokens, err := db.FetchAPITokens(c.Request().Context(), DefaultPageNumber, DefaultPageSize)
		if err != nil {
			c.Logger().Errorf("db.FetchAPITokens failed with err: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load api tokens")
		}

		presentedTokens := []presenter.APIToken{}

		for _, apiToken := range apiTokens {
			presentedTokens = append(presentedTokens, presenter.APITokenFromModel(apiToken))
		}

		pageHTML := view.APITokens(presentedTokens)
		htmlString := view.Layout("fundock | api tokens", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
