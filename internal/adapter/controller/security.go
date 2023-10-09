package controller

import (
	"context"
	"log/slog"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/pkg/api"
	"github.com/otakakot/ninshow/pkg/log"
)

var _ api.SecurityHandler = (*Security)(nil)

type Security struct {
	security usecase.Security
}

func NewSecurity(
	security usecase.Security,
) *Security {
	return &Security{
		security: security,
	}
}

// HandleBearer implements api.SecurityHandler.
func (sec *Security) HandleBearer(
	ctx context.Context,
	operationName string,
	t api.Bearer,
) (context.Context, error) {
	end := log.StartEnd(ctx)
	defer end()

	slog.InfoContext(ctx, operationName)

	output, err := sec.security.HandleBearer(ctx, usecase.HandleBearerInput{
		Token: t.Token,
	})
	if err != nil {
		return ctx, err
	}

	return output.Ctx, nil
}
