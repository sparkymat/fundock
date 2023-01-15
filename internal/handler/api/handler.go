package api

import "github.com/labstack/echo/v4"

func renderError(c echo.Context, statusCode int, message string) error {
	//nolint:wrapcheck
	return c.JSON(statusCode, map[string]string{
		"error": message,
	})
}
