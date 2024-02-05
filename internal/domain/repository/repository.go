package repository

import (
	"context"
	"time"

	"github.com/otakakot/ninshow/internal/domain/model"
)

type Account interface {
	Save(ctx context.Context, account model.Account) error
	Find(ctx context.Context, id string) (*model.Account, error)
	FindByEmail(ctx context.Context, email string) (*model.Account, error)
	FindPassword(ctx context.Context, id string) ([]byte, error)
}

type OIDCClient interface {
	Save(ctx context.Context, client model.OIDCClient) error
	Find(ctx context.Context, id string) (*model.OIDCClient, error)
	FindSecret(ctx context.Context, id string) ([]byte, error)
}

type JWTSignKey interface {
	Save(ctx context.Context, key model.JWTSignKey) error
	Find(ctx context.Context, id string) (*model.JWTSignKey, error)
	List(ctx context.Context) ([]model.JWTSignKey, error)
}

type Cache[T any] interface {
	Set(ctx context.Context, key string, value T, ttl time.Duration) error
	Get(ctx context.Context, key string) (T, error)
	Del(ctx context.Context, key string) error
}
