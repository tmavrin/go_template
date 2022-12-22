package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/teltech/logger"
	"github.com/tmavrin/go_template/config"
	v1 "github.com/tmavrin/go_template/internal/http/v1"
	"github.com/tmavrin/go_template/internal/repository/postgres"
	"github.com/tmavrin/go_template/internal/service"
)

type (
	server struct {
		httpServer *http.Server
		Close      func()
		log        *logger.Log
		services   APIServices
	}

	APIServices struct {
		AccountService service.AccountService
	}
)

func NewServer(ctx context.Context, log *logger.Log) (*server, error) {
	var env config.APIEnvironment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, fmt.Errorf("initialize api environment variables: %w", err)
	}

	db, err := postgres.New(ctx, env.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed create database connection: %w", err)
	}

	router := v1.Routes(&v1.Resources{
		Log:            log,
		AccountManager: postgres.NewAccountManager(db, log),
	})

	s := &server{
		httpServer: &http.Server{
			Addr:    ":" + env.HTTPPort,
			Handler: router,
		},
		Close: func() {
			log.Info("closing gracefully database and stackdriver")
		},
		log: log,

		services: APIServices{
			AccountService: service.AccountService{
				AccountProvider: postgres.NewAccountManager(db, log),
			},
		},
	}

	return s, nil
}

func (s *server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
