package service

import (
	"context"

	"github.com/teltech/logger"
	"github.com/tmavrin/go_template/internal"
)

type AccountService struct {
	AccountProvider internal.AccountProvider
	Log             *logger.Log
}

func (as *AccountService) CreateAccount(ctx context.Context, account internal.Account) (internal.Account, error) {

	account, err := as.AccountProvider.Create(ctx, account)

	// do something else

	// send sms verification

	// LP events

	return account, err
}
