package controller

import (
	"context"
	"log/slog"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/pkg/api"
)

var _ api.Handler = (*Controller)(nil)

type Controller struct {
	account usecase.Account
}

func NewController(
	account usecase.Account,
) *Controller {
	return &Controller{
		account: account,
	}
}

// Health implements api.Handler.
func (*Controller) Health(ctx context.Context) (api.HealthRes, error) {
	slog.Info("start health controller")
	defer slog.Info("end health controller")
	return &api.HealthOK{}, nil
}

// IdpSignup implements api.Handler.
func (ctl *Controller) IdpSignup(
	ctx context.Context,
	req api.OptIdPSignupRequestSchema,
) (api.IdpSignupRes, error) {
	slog.Info("start idp signup controller")
	defer slog.Info("end idp signup controller")

	if _, err := ctl.account.Signup(ctx, usecase.AccountSignupInput{
		Email:    req.Value.Email,
		Username: req.Value.Username,
		Password: req.Value.Password,
	}); err != nil {
		return &api.IdpSignupInternalServerError{}, nil
	}

	return &api.IdpSignupOK{}, nil
}

// IdpSignin implements api.Handler.
func (ctl *Controller) IdpSignin(
	ctx context.Context,
	req api.OptIdPSigninRequestSchema,
) (api.IdpSigninRes, error) {
	slog.Info("start idp signin controller")
	defer slog.Info("end idp signin controller")

	if _, err := ctl.account.Signin(ctx, usecase.AccountSigninInput{
		Username: req.Value.Username,
		Password: req.Value.Password,
	}); err != nil {
		return &api.IdpSigninUnauthorized{}, nil
	}

	return &api.IdpSigninOK{}, nil
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
