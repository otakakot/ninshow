package interactor

import (
	"context"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
)

var _ usecase.Security = (*Security)(nil)

type Security struct {
	cache repository.Cache[struct{}]
}

func NewSecurity(
	cache repository.Cache[struct{}],
) *Security {
	return &Security{
		cache: cache,
	}
}

// HandleBearer implements usecase.Security.
func (sec *Security) HandleBearer(
	ctx context.Context,
	input usecase.HandleBearerInput,
) (*usecase.HandleBearerOutput, error) {
	// Access Token が有効か確認する
	if _, err := model.ParseAccessToken(input.Token, input.Sign); err != nil {
		return nil, err
	}

	// Access Token が Cache に存在するか確認する
	if _, err := sec.cache.Get(ctx, input.Token); err != nil {
		return nil, err
	}

	// TODO: context に必要な値を保存する

	return &usecase.HandleBearerOutput{
		Ctx: ctx,
	}, nil
}
