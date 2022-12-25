package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config"
	"github.com/sparkymat/fundock/database"
	"github.com/sparkymat/fundock/route"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := database.New(database.Config{
		ConnectionString: cfg.DBConnectionString(),
	})
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(); err != nil {
		panic(err)
	}

	e := echo.New()
	route.Setup(e, cfg, db)

	e.Logger.Fatal(e.Start(":8080"))
}
