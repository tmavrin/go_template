package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teltech/logger"
	"github.com/tmavrin/go_template/config"
	"github.com/tmavrin/go_template/internal/repository/postgres"
	"github.com/tmavrin/go_template/internal/service"
)

type Resources struct {
	Log            *logger.Log
	Environment    config.APIEnvironment
	AccountManager *postgres.AccountManager
}

func Routes(resource *Resources) http.Handler {
	r := chi.NewRouter()

	accountRouter := NewAccountRouter(resource.Log,
		service.AccountService{
			Log:             resource.Log,
			AccountProvider: resource.AccountManager,
		})

	r.Mount("/account", accountRouter.routes())

	return r
}
