package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teltech/logger"
	"github.com/tmavrin/go_template/internal"
)

type AccountManager struct {
	db  *pgxpool.Pool
	log *logger.Log
}

func NewAccountManager(conn *pgxpool.Pool, log *logger.Log) *AccountManager {
	return &AccountManager{
		db:  conn,
		log: log,
	}
}

func (am *AccountManager) Create(ctx context.Context, Account internal.Account) (internal.Account, error) {
	return internal.Account{
		Name: "Test User",
	}, nil
}
