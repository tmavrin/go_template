package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"github.com/teltech/logger"
	"github.com/tmavrin/go_template/internal"
	"github.com/tmavrin/go_template/internal/http/httphelper"
	"github.com/tmavrin/go_template/internal/service"
)

type (
	AccountCreateDTO struct {
		Email       null.String `json:"email"`
		PhoneNumber null.String `json:"phone_number"`
		Name        string      `json:"name"`
	}

	accountRouter struct {
		log            *logger.Log
		accountService service.AccountService
	}
)

func NewAccountRouter(log *logger.Log, accountService service.AccountService) *accountRouter {
	return &accountRouter{
		log:            log,
		accountService: accountService,
	}
}

func (ar *accountRouter) routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", ar.handleCreateAccount)
	r.Get("/", ar.handleGetAccount)
	r.Put("/", ar.handleUpdateAccount)
	r.Delete("/", ar.handleDeleteAccount)

	return r
}

func (ar *accountRouter) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var (
		account AccountCreateDTO

		log = ar.log.With(logger.Fields{
			"handler": "accountRouter.handleCreateAccount",
		})
	)

	if err := httphelper.DecodeJSONBody(r, &account); err != nil {
		log.Errorf("failed to decode account: %s", err)
		httphelper.JSONErrorFromHTTPStatus(log, w, http.StatusBadRequest)
		return
	}

	result, err := ar.accountService.CreateAccount(r.Context(),
		internal.Account{
			Email:       account.Email,
			PhoneNumber: account.PhoneNumber,
			Name:        account.Name,
		},
	)

	if err != nil {
		log.Errorf("failed to create account: %s", err)
		httphelper.JSONErrorFromHTTPStatus(log, w, http.StatusInternalServerError)
		return
	}

	if err := httphelper.WriteJSONResponse(w, result); err != nil {
		log.Errorf("failed to encode json response: %s", err)
		httphelper.JSONErrorFromHTTPStatus(log, w, http.StatusInternalServerError)
		return
	}

}

func (ar *accountRouter) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (ar *accountRouter) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (ar *accountRouter) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	// TODO
}
