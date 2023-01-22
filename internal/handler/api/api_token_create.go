package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/internal/handler/api/presenter"
)

type APITokenCreateInput struct {
	ClientName string `json:"client_name"` //nolint:tagliatelle
}

func APITokenCreate(_ configiface.ConfigAPI, db dbiface.DBAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &APITokenCreateInput{}
		if err := c.Bind(input); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "invalid input")
		}

		token := strings.ToLower(ulid.Make().String())

		id, err := db.CreateAPIToken(c.Request().Context(), input.ClientName, token)
		if err != nil {
			c.Logger().Errorf("db.CreateAPIToken failed with err: %v", err)

			return renderError(c, http.StatusInternalServerError, "token creation failed")
		}

		apiToken, err := db.FetchAPITokenByID(c.Request().Context(), *id)
		if err != nil {
			c.Logger().Errorf("db.FetchAPIToken failed with err: %v", err)

			return renderError(c, http.StatusInternalServerError, "api token fetch failed")
		}

		presentedToken := presenter.APITokenFromModel(*apiToken)

		//nolint:wrapcheck
		return c.JSON(http.StatusOK, presentedToken)
	}
}
