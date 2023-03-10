package route

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
	"github.com/sparkymat/fundock/view"
)

func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := err.Error()

	//nolint:errorlint
	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
		message = fmt.Sprintf("%v", httpErr.Message)
	}

	c.Logger().Error(err)

	pageHTML := view.Error(fmt.Sprintf("%v", message))
	htmlString := view.BasicLayout("fundock | error", pageHTML)

	if err := c.HTMLBlob(code, []byte(htmlString)); err != nil {
		c.Logger().Error(err)
	}
}

func Setup(e *echo.Echo, cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) {
	e.HTTPErrorHandler = customErrorHandler
	e.Use(middleware.Recover())
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/fonts", "public/fonts")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret()))))

	setupWebRoutes(e, cfg, db)
	setupAPIRoutes(e, cfg, db, dockerSvc)
}
