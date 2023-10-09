package controller

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/pkg/api"
	"github.com/otakakot/ninshow/pkg/log"
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
	ctx context.Context,
) (api.HealthRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	return &api.HealthOK{}, nil
}

// IdpSignup implements api.Handler.
func (ctl *Controller) IdpSignup(
	ctx context.Context,
	req api.OptIdPSignupRequestSchema,
) (api.IdpSignupRes, error) {
	end := log.StartEnd(ctx)
	defer end()

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
	end := log.StartEnd(ctx)
	defer end()

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
	end := log.StartEnd(ctx)
	defer end()

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
func (ctl *Controller) OpCallback(
	ctx context.Context,
	params api.OpCallbackParams,
) (api.OpCallbackRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	output, err := ctl.op.Callback(ctx, usecase.OpenIDProviderCallbackInput{
		ID: params.ID,
	})
	if err != nil {
		return &api.OpCallbackInternalServerError{}, err
	}

	res := &api.OpCallbackFound{}

	res.SetLocation(api.NewOptURI(output.RedirectURI))

	return res, nil
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
	end := log.StartEnd(ctx)
	defer end()

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
	end := log.StartEnd(ctx)
	defer end()

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
func (ctl *Controller) OpOpenIDConfiguration(
	ctx context.Context,
) (api.OpOpenIDConfigurationRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	output, err := ctl.op.Configuration(ctx, usecase.OpenIDProviderConfigurationInput{})
	if err != nil {
		return &api.OpOpenIDConfigurationInternalServerError{}, err
	}

	return &api.OPOpenIDConfigurationResponseSchema{
		Issuer:                output.Issuer,
		AuthorizationEndpoint: output.AuthorizationEndpoint,
		TokenEndpoint:         output.TokenEndpoint,
		UserinfoEndpoint:      output.UserinfoEndpoint,
		JwksURL:               output.JwksURL,
		RevocationEndpoint:    output.RevocationEndpoint,
	}, nil
}

// OpRevoke implements api.Handler.
func (*Controller) OpRevoke(ctx context.Context, req *api.OPRevokeRequestSchema) (api.OpRevokeRes, error) {
	panic("unimplemented")
}

// OpToken implements api.Handler.
func (*Controller) OpToken(
	ctx context.Context,
	req *api.OPTokenRequestSchema,
) (api.OpTokenRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	res := &api.OPTokenResponseSchemaHeaders{
		CacheControl: api.NewOptString("no-store"),
		Pragma:       api.NewOptString("no-cache"),
		Response:     api.OPTokenResponseSchema{},
	}

	return res, nil
}

// OpUserinfo implements api.Handler.
func (*Controller) OpUserinfo(ctx context.Context) (api.OpUserinfoRes, error) {
	panic("unimplemented")
}

// RpCallback implements api.Handler.
func (*Controller) RpCallback(
	ctx context.Context,
	params api.RpCallbackParams,
) (api.RpCallbackRes, error) {
	panic("unimplemented")
}

// RpLogin implements api.Handler.
func (ctl *Controller) RpLogin(
	ctx context.Context,
) (api.RpLoginRes, error) {
	end := log.StartEnd(ctx)
	defer end()

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
