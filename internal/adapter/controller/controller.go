package controller

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/otakakot/ninshow/internal/application/usecase"
	derror "github.com/otakakot/ninshow/internal/domain/errors"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/pkg/api"
	"github.com/otakakot/ninshow/pkg/log"
)

type Config interface {
	SelfEndpoint() string
	OIDCEndpoint() string
	IDTokenSignKey() *rsa.PrivateKey
	AcessTokenSign() string
	RelyingPartyID() string
	RelyingPartySecret() string
}

var _ api.Handler = (*Controller)(nil)

type Controller struct {
	config Config
	idp    usecase.IdentityProvider
	op     usecase.OpenIDProviider
	rp     usecase.RelyingParty
}

func NewController(
	config Config,
	idp usecase.IdentityProvider,
	op usecase.OpenIDProviider,
	rp usecase.RelyingParty,
) *Controller {
	return &Controller{
		config: config,
		idp:    idp,
		op:     op,
		rp:     rp,
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

	slog.Info("signup", "email", req.Value.Email, "password", req.Value.Password)

	if _, err := ctl.idp.Signup(ctx, usecase.IdentityProviderSignupInput{
		Email:    req.Value.Email,
		Name:     req.Value.Name,
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
		Email:    req.Value.Email,
		Password: req.Value.Password,
	}); err != nil {
		slog.WarnContext(ctx, err.Error())

		return &api.IdpSigninUnauthorized{}, nil
	}

	return &api.IdpSigninOK{}, nil
}

// IdpOIDCLogin implements api.Handler.
func (clt *Controller) IdpOIDCLogin(
	ctx context.Context,
	params api.IdpOIDCLoginParams,
) (api.IdpOIDCLoginRes, error) {
	panic("unimplemented")
}

// IdpOIDCCallback implements api.Handler.
func (ctl *Controller) IdpOIDCCallback(
	ctx context.Context,
	params api.IdpOIDCCallbackParams,
) (api.IdpOIDCCallbackRes, error) {
	panic("unimplemented")
}

// OpAuthorize implements api.Handler.
func (ctl *Controller) OpAuthorize(
	ctx context.Context,
	params api.OpAuthorizeParams,
) (api.OpAuthorizeRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	output, err := ctl.op.Authorize(ctx, usecase.OpenIDProviderAuthorizeInput{
		LoginURL:     fmt.Sprintf("%s/op/login", ctl.config.SelfEndpoint()),
		ResponseType: string(params.ResponseType),
		Scope:        strings.Split(params.Scope, " "),
		ClientID:     params.ClientID.String(),
		RedirectURI:  params.RedirectURI.String(),
		State:        params.State.Value,
		Nonce:        params.Nonce.Value,
	})
	if err != nil {
		slog.Warn(err.Error())

		switch {
		case errors.Is(err, derror.ErrInvalidRequest):
			return &api.OpAuthorizeBadRequest{
				Error: api.NewOptOpAuthorizeBadRequestError(api.OpAuthorizeBadRequestErrorInvalidRequest),
			}, nil
		case errors.Is(err, derror.ErrUnauthorizedClient):
			return &api.OpAuthorizeUnauthorized{
				Error: api.NewOptOpAuthorizeUnauthorizedError(api.OpAuthorizeUnauthorizedErrorUnauthorizedClient),
			}, nil
		case errors.Is(err, derror.ErrAccessDenied):
			return &api.OpAuthorizeForbidden{
				Error: api.NewOptOpAuthorizeForbiddenError(api.OpAuthorizeForbiddenErrorAccessDenied),
			}, nil
		case errors.Is(err, derror.ErrInvalidScope):
			return &api.OpAuthorizeBadRequest{
				Error: api.NewOptOpAuthorizeBadRequestError(api.OpAuthorizeBadRequestErrorInvalidScope),
			}, nil
		default:
			return &api.OpAuthorizeInternalServerError{
				Error: api.NewOptOpAuthorizeInternalServerErrorError(api.OpAuthorizeInternalServerErrorErrorServerError),
			}, err
		}
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
		ID:             params.ID,
		Issuer:         ctl.config.SelfEndpoint(),
		IDTokenSignKey: ctl.config.IDTokenSignKey(),
	})
	if err != nil {
		return &api.OpCallbackInternalServerError{}, err
	}

	res := &api.OpCallbackFound{}

	res.SetLocation(api.NewOptURI(output.RedirectURI))

	return res, nil
}

// OpCerts implements api.Handler.
func (ctl *Controller) OpCerts(
	ctx context.Context,
) (api.OpCertsRes, error) {
	output, err := ctl.op.Certs(ctx, usecase.OpenIDProviderCertsInput{
		PublicKey: ctl.config.IDTokenSignKey().PublicKey,
	})
	if err != nil {
		return &api.OpCertsInternalServerError{}, err
	}

	return &api.OPJWKSetResponseSchema{
		Keys: []api.OPJWKSetKey{
			{
				Kid: output.Kid,
				Kty: output.Kty,
				Use: output.Use,
				Alg: output.Alg,
				N:   output.N,
				E:   output.E,
			},
		},
	}, nil
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
		Email:       req.Email,
		Password:    req.Password,
		CallbackURL: fmt.Sprintf("%s/op/callback", ctl.config.SelfEndpoint()),
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

	slog.Info("login view", "auth_request_id", params.AuthRequestID)

	return &api.OpLoginViewOKHeaders{
		XRequestID: api.NewOptString(params.AuthRequestID),
		Response: api.OpLoginViewOK{
			Data: output.Data,
		},
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
		JwksURI:               output.JwksURI,
		RevocationEndpoint:    output.RevocationEndpoint,
	}, nil
}

// OpToken implements api.Handler.
func (ctl *Controller) OpToken(
	ctx context.Context,
	req *api.OPTokenRequestSchema,
) (api.OpTokenRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	switch req.GrantType {
	case api.OPTokenRequestSchemaGrantTypeAuthorizationCode:
		output, err := ctl.op.AuthorizationCodeGrant(ctx, usecase.OpenIDProviderAuthorizationCodeGrantInput{
			ClientID:        req.ClientID.Value,
			ClientSecret:    req.ClientSecret.Value,
			Issuer:          ctl.config.SelfEndpoint(),
			Code:            req.Code,
			AccessTokenSign: ctl.config.AcessTokenSign(),
			IDTokenSignKey:  ctl.config.IDTokenSignKey(),
		})
		if err != nil {
			return &api.OpTokenInternalServerError{}, err
		}

		scope := make([]api.OPTokenResponseSchemaScopeItem, 0, 3)
		for _, v := range strings.Split(req.Scope.Value, " ") {
			scope = append(scope, api.OPTokenResponseSchemaScopeItem(v))
		}

		res := &api.OPTokenResponseSchemaHeaders{
			CacheControl: api.NewOptString("no-store"),
			Pragma:       api.NewOptString("no-cache"),
			Response: api.OPTokenResponseSchema{
				AccessToken:  output.AccessToken,
				TokenType:    output.TokenType,
				RefreshToken: output.RefreshToken,
				ExpiresIn:    output.ExpiresIn,
				IDToken:      output.IDToken,
				Scope:        scope,
			},
		}

		return res, nil
	case api.OPTokenRequestSchemaGrantTypeRefreshToken:
		scope := append(make([]string, 0, 3), strings.Split(req.Scope.Value, " ")...)

		output, err := ctl.op.RefreshTkenGrant(ctx, usecase.OpenIDProviderRefreshTokenGrantInput{
			RefreshToken:    req.RefreshToken.Value,
			ClientID:        req.ClientID.Value,
			Scope:           scope,
			Issuer:          ctl.config.SelfEndpoint(),
			AccessTokenSign: ctl.config.AcessTokenSign(),
			IDTokenSignKey:  ctl.config.IDTokenSignKey(),
		})
		if err != nil {
			return &api.OpTokenInternalServerError{}, err
		}

		sc := make([]api.OPTokenResponseSchemaScopeItem, len(output.Scope))
		for i, v := range output.Scope {
			sc[i] = api.OPTokenResponseSchemaScopeItem(v)
		}

		res := &api.OPTokenResponseSchemaHeaders{
			CacheControl: api.NewOptString("no-store"),
			Pragma:       api.NewOptString("no-cache"),
			Response: api.OPTokenResponseSchema{
				AccessToken:  output.AccessToken,
				TokenType:    output.TokenType,
				RefreshToken: output.RefreshToken,
				ExpiresIn:    output.ExpiresIn,
				IDToken:      output.IDToken,
				Scope:        sc,
			},
		}

		return res, nil
	default:
		return &api.OpTokenBadRequest{
			Error: api.NewOptString(fmt.Errorf("invalid grant_type: %s", req.GrantType).Error()),
		}, nil
	}
}

// OpUserinfo implements api.Handler.
func (ctl *Controller) OpUserinfo(
	ctx context.Context,
) (api.OpUserinfoRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	at := model.GetAccessTokenCtx(ctx)

	output, err := ctl.op.Userinfo(ctx, usecase.OpenIDProviderUserinfoInput{
		AccessToken: at,
	})
	if err != nil {
		return &api.OpUserinfoInternalServerError{}, err
	}

	res := &api.OPUserInfoResponseSchema{
		Sub: output.Sub,
	}

	if output.Profile != nil {
		res.Profile = api.NewOptString(*output.Profile)
	}

	if output.Email != nil {
		res.Email = api.NewOptString(*output.Email)
	}

	return res, nil
}

// OpRevoke implements api.Handler.
func (ctl *Controller) OpRevoke(ctx context.Context, req *api.OPRevokeRequestSchema) (api.OpRevokeRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	if _, err := ctl.op.Revoke(ctx, usecase.OpenIDProviderRevokeInput{
		Hint:  string(req.TokenTypeHint.Value),
		Token: req.Token,
	}); err != nil {
		return &api.OpRevokeInternalServerError{}, err
	}

	return &api.OpRevokeOK{}, nil
}

// RpCallback implements api.Handler.
func (ctl *Controller) RpCallback(
	ctx context.Context,
	params api.RpCallbackParams,
) (api.RpCallbackRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	// CSRF 対策
	if params.QueryState != params.CookieState {
		return &api.RpCallbackBadRequest{}, nil
	}

	output, err := ctl.rp.Callback(ctx, usecase.RelyingPartyCallbackInput{
		Code:         params.Code,
		OIDCEndpoint: ctl.config.SelfEndpoint(),
		ClientID:     ctl.config.RelyingPartyID(),
		ClientSecret: ctl.config.RelyingPartySecret(),
	})
	if err != nil {
		return &api.RpCallbackInternalServerError{}, err
	}

	return &api.RpCallbackOK{
		Data: output.Data,
	}, nil
}

// RpLogin implements api.Handler.
func (ctl *Controller) RpLogin(
	ctx context.Context,
) (api.RpLoginRes, error) {
	end := log.StartEnd(ctx)
	defer end()

	output, err := ctl.rp.Login(ctx, usecase.RelyingPartyLoginInput{
		OIDCEndpoint: fmt.Sprintf("%s/authorize", ctl.config.OIDCEndpoint()),
		ClientID:     ctl.config.RelyingPartyID(),
		RedirectURI:  fmt.Sprintf("%s/rp/callback", ctl.config.SelfEndpoint()),
		Scope:        []string{"openid", "profile", "email"},
	})
	if err != nil {
		return &api.RpLoginInternalServerError{}, err
	}

	res := &api.RpLoginFound{}

	res.SetSetCookie(api.NewOptString(output.Cookie.String()))

	res.SetLocation(api.NewOptURI(output.RedirectURI))

	return res, nil
}
