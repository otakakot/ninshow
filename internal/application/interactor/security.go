package interactor

import (
	"context"
	"fmt"

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
	at, err := model.ParseAccessToken(input.Token, input.Sign)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	// Access Token が Cache に存在するか確認する
	if _, err := sec.cache.Get(ctx, input.Token); err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	ctx = model.SetAccessTokenCtx(ctx, at)

	return &usecase.HandleBearerOutput{
		Ctx: ctx,
	}, nil
}
