package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"math/big"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/fundock/config"
	"github.com/sparkymat/fundock/database"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker"
	"github.com/sparkymat/fundock/internal/route"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	dockerSvc, err := docker.New()
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

	if cfg.SingleUser() {
		err = ensureAdminUser(db)
		if err != nil {
			panic(err)
		}
	}

	e := echo.New()
	route.Setup(e, cfg, db, dockerSvc)

	e.Logger.Fatal(e.Start(":8080"))
}

func ensureAdminUser(db dbiface.DBAPI) error {
	_, err := db.FetchUser(context.Background(), "admin")
	if err == nil {
		return nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	password, err := generateRandomPassword()
	if err != nil {
		return err
	}

	encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	_, err = db.CreateUser(
		context.Background(),
		"admin",
		string(encryptedPasswordBytes),
		nil,
		nil,
	)

	return err
}

func generateRandomPassword() (string, error) {
	password := ""

	characters := "abcdefghijklmnopqrstuvwxyz0123456789"

	charsLeft := 16

	for charsLeft > 0 {
		position, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}

		posInt := int(position.Int64())
		password = fmt.Sprintf("%s%c", password, characters[posInt])
		charsLeft--
	}

	return password, nil
}
