package repository

import (
	"context"
	"time"

	"github.com/otakakot/ninshow/internal/domain/model"
)

type Account interface {
	Save(context.Context, model.Account) error
	Find(context.Context, string) (*model.Account, error)
}

type Cache[T any] interface {
	Set(context.Context, string, T, time.Duration) error
	Get(context.Context, string) (T, error)
	Del(context.Context, string) error
}
