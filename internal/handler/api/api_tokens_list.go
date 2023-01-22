package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
)

type APITokensListInput struct {
	PageSize   uint32 `query:"page_size"`
	PageNumber uint32 `query:"page_number"`
}

func APITokensList(_ configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &APITokensListInput{}
		if err := c.Bind(input); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "invalid input")
		}

		apiTokens, err := db.FetchAPITokens(c.Request().Context(), input.PageNumber, input.PageSize)
		if err != nil {
			c.Logger().Errorf("db.FetchAPITokens failed with err: %v", err)

			return renderError(c, http.StatusInternalServerError, "failed to load api tokens")
		}

		presentedTokensList := presenter.APITokensListFromModels(input.PageNumber, input.PageSize, apiTokens)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedTokensList)
	}
}
