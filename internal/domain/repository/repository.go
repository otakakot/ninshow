package repository

import (
	"context"
	"time"

	"github.com/otakakot/ninshow/internal/domain/model"
)

type Account interface {
	Save(context.Context, model.Account) error
	Find(context.Context, string) (*model.Account, error)
	FindByEmail(context.Context, string) (*model.Account, error)
	FindPassword(context.Context, string) ([]byte, error)
}

type OIDCClient interface {
	Save(context.Context, model.OIDCClient) error
	Find(context.Context, string) (*model.OIDCClient, error)
	FindSecret(context.Context, string) ([]byte, error)
}

type Cache[T any] interface {
	Set(context.Context, string, T, time.Duration) error
	Get(context.Context, string) (T, error)
	Del(context.Context, string) error
}
