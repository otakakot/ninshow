package controller

import (
	"context"

	"github.com/otakakot/ninshow/pkg/api"
)

var _ api.Handler = (*Controller)(nil)

type Controller struct {
}

// Health implements api.Handler.
func (*Controller) Health(ctx context.Context) (api.HealthRes, error) {
	return &api.HealthOK{}, nil
}

// OpAuthorize implements api.Handler.
func (*Controller) OpAuthorize(ctx context.Context, params api.OpAuthorizeParams) (api.OpAuthorizeRes, error) {
	panic("unimplemented")
}

// OpCallback implements api.Handler.
func (*Controller) OpCallback(ctx context.Context, params api.OpCallbackParams) (api.OpCallbackRes, error) {
	panic("unimplemented")
}

// OpCerts implements api.Handler.
func (*Controller) OpCerts(ctx context.Context) (api.OpCertsRes, error) {
	panic("unimplemented")
}

// OpLogin implements api.Handler.
func (*Controller) OpLogin(ctx context.Context) (api.OpLoginRes, error) {
	panic("unimplemented")
}

// OpLoginView implements api.Handler.
func (*Controller) OpLoginView(ctx context.Context, params api.OpLoginViewParams) (api.OpLoginViewRes, error) {
	panic("unimplemented")
}

// OpOpenIDConfiguration implements api.Handler.
func (*Controller) OpOpenIDConfiguration(ctx context.Context) (api.OpOpenIDConfigurationRes, error) {
	panic("unimplemented")
}

// OpRevoke implements api.Handler.
func (*Controller) OpRevoke(ctx context.Context, req *api.OPRevokeRequestSchema) (api.OpRevokeRes, error) {
	panic("unimplemented")
}

// OpToken implements api.Handler.
func (*Controller) OpToken(ctx context.Context, req *api.OPTokenRequestSchema) (api.OpTokenRes, error) {
	panic("unimplemented")
}

// OpUserinfo implements api.Handler.
func (*Controller) OpUserinfo(ctx context.Context) (api.OpUserinfoRes, error) {
	panic("unimplemented")
}

// RpCallback implements api.Handler.
func (*Controller) RpCallback(ctx context.Context, params api.RpCallbackParams) (api.RpCallbackRes, error) {
	panic("unimplemented")
}

// RpLogin implements api.Handler.
func (*Controller) RpLogin(ctx context.Context) (api.RpLoginRes, error) {
	panic("unimplemented")
}
