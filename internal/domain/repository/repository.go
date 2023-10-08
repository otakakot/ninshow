package repository

import (
	"context"

	"github.com/otakakot/ninshow/internal/domain/model"
)

type Account interface {
	Save(context.Context, model.Account) error
	Find(context.Context, string) (*model.Account, error)
}
