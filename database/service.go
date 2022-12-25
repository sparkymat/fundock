package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Importing file driver for migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	ConnectionString string
}

func New(cfg Config) (*Service, error) {
	dbConn, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return &Service{
		conn: dbConn,
	}, nil
}

type Service struct {
	conn *sqlx.DB
}

func (s *Service) AutoMigrate() error {
	driver, err := postgres.WithInstance(s.conn.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver. err: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration driver. err: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations. err: %w", err)
	}

	return nil
}
