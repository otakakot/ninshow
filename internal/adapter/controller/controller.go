package controller

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/pkg/api"
)

var _ api.Handler = (*Controller)(nil)

type Controller struct {
	idp usecase.IdentityProvider
	op  usecase.OpenIDProviider
	rp  usecase.RelyingParty
}

func NewController(
	idp usecase.IdentityProvider,
	op usecase.OpenIDProviider,
	rp usecase.RelyingParty,
) *Controller {
	return &Controller{
		idp: idp,
		op:  op,
		rp:  rp,
	}
}

// Health implements api.Handler.
func (*Controller) Health(
	_ context.Context,
) (api.HealthRes, error) {
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

	slog.Info("signup", "username", req.Value.Username, "password", req.Value.Password)

	if _, err := ctl.idp.Signup(ctx, usecase.IdentityProviderSignupInput{
		Email:    req.Value.Email,
		Username: req.Value.Username,
		Password: req.Value.Password,
	}); err != nil {
		slog.WarnContext(ctx, err.Error())

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

	if _, err := ctl.idp.Signin(ctx, usecase.IdentityProviderSigninInput{
		Username: req.Value.Username,
		Password: req.Value.Password,
	}); err != nil {
		slog.WarnContext(ctx, err.Error())

		return &api.IdpSigninUnauthorized{}, nil
	}

	return &api.IdpSigninOK{}, nil
}

// OpAuthorize implements api.Handler.
func (ctl *Controller) OpAuthorize(
	ctx context.Context,
	params api.OpAuthorizeParams,
) (api.OpAuthorizeRes, error) {
	slog.Info("start op authorize controller")
	defer slog.Info("end op authorize controller")

	scope := make([]string, len(params.Scope))
	for i, v := range params.Scope {
		scope[i] = string(v)
	}

	output, err := ctl.op.Autorize(ctx, usecase.OpenIDProviderAuthorizeInput{
		LoginURL:     fmt.Sprintf("%s/op/login", "http://localhost:8080"),
		ResponseType: string(params.ResponseType),
		Scope:        scope,
		ClientID:     params.ClientID.String(),
		RedirectURI:  params.RedirectURI.String(),
		State:        params.State.Value,
		Nonce:        params.Nonce.Value,
	})
	if err != nil {
		return &api.OpAuthorizeInternalServerError{}, err
	}

	res := &api.OpAuthorizeFound{}

	res.SetLocation(api.NewOptURI(output.RedirectURI))

	return res, nil
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
func (ctl *Controller) OpLogin(
	ctx context.Context,
	req *api.OPLoginRequestSchema,
) (api.OpLoginRes, error) {
	slog.Info("start op login controller")
	defer slog.Info("end op login controller")

	output, err := ctl.op.Login(ctx, usecase.OpenIDProviderLoginInput{
		ID:          req.ID,
		Username:    req.Username,
		Password:    req.Password,
		CallbackURL: "http://localhost:8080/op/callback",
	})
	if err != nil {
		return &api.OpLoginInternalServerError{}, err
	}

	res := &api.OpLoginFound{}

	res.SetLocation(api.NewOptURI(output.RedirectURI))

	return res, nil
}

// OpLoginView implements api.Handler.
func (ctl *Controller) OpLoginView(ctx context.Context, params api.OpLoginViewParams) (api.OpLoginViewRes, error) {
	slog.Info("start op login view controller")
	defer slog.Info("end op login view controller")

	output, err := ctl.op.LoginVeiw(ctx, usecase.OpenIDProviderLoginViewInput{
		AuthRequestID: params.AuthRequestID,
	})
	if err != nil {
		return &api.OpLoginViewInternalServerError{}, err
	}

	return &api.OpLoginViewOK{
		Data: output.Data,
	}, nil
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
func (ctl *Controller) RpLogin(
	ctx context.Context,
) (api.RpLoginRes, error) {
	slog.Info("start rp login controller")
	defer slog.Info("end rp login controller")

	output, err := ctl.rp.Login(ctx, usecase.RelyingPartyLoginInput{
		OIDCEndpoint: "http://localhost:8080/op/authorize",
		ClientID:     "test",
		RedirectURI:  "http://localhost:8080/rp/callback",
		Scope:        []string{string(api.OpAuthorizeScopeItemOpenid)},
	})
	if err != nil {
		return &api.RpLoginInternalServerError{}, err
	}

	res := &api.RpLoginFound{}

	res.SetSetCookie(api.NewOptString(output.Cookie.String()))

	res.SetLocation(api.NewOptURI(output.RedirectURI))

	return res, nil
}
