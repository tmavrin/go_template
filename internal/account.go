package internal

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Account struct {
	ID           uuid.UUID   `json:"id"`
	ThirdPartyID string      `json:"third_party_id"`
	Email        null.String `json:"email"`
	PhoneNumber  null.String `json:"phone_number"`
	Name         string      `json:"name"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type AccountProvider interface {
	Create(ctx context.Context, account Account) (Account, error)
}
