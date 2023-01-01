package runner

import (
	"errors"

	"github.com/sparkymat/fundock/config/configiface"
	"github.com/sparkymat/fundock/database/dbiface"
	"github.com/sparkymat/fundock/docker/dockeriface"
)

var ErrDBFailure = errors.New("db failure")

func New(cfg configiface.ConfigAPI, db dbiface.DBAPI, dockerSvc dockeriface.DockerAPI) (*Service, error) {
	return &Service{
		cfg:       cfg,
		db:        db,
		dockerSvc: dockerSvc,
	}, nil
}

type Service struct {
	cfg       configiface.ConfigAPI
	db        dbiface.DBAPI
	dockerSvc dockeriface.DockerAPI
}
