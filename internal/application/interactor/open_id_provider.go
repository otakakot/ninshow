package interactor

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"html/template"
	"log/slog"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/otakakot/ninshow/internal/application/usecase"
	derror "github.com/otakakot/ninshow/internal/domain/errors"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
	"github.com/otakakot/ninshow/pkg/log"
)

var _ usecase.OpenIDProviider = (*OpenIDProvider)(nil)

type OpenIDProvider struct {
	client       repository.OIDCClient
	account      repository.Account
	param        repository.Cache[model.AuthorizeParam]
	loggedin     repository.Cache[model.LoggedIn]
	accessToken  repository.Cache[struct{}]
	refreshToken repository.Cache[string]
}

func NewOpenIDProvider(
	client repository.OIDCClient,
	account repository.Account,
	param repository.Cache[model.AuthorizeParam],
	loggedin repository.Cache[model.LoggedIn],
	accessToken repository.Cache[struct{}],
	refreshToken repository.Cache[string],
) *OpenIDProvider {
	return &OpenIDProvider{
		client:       client,
		account:      account,
		loggedin:     loggedin,
		param:        param,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

// Configuration implements usecase.OpenIDProviider.
func (*OpenIDProvider) Configuration(
	ctx context.Context, input usecase.OpenIDProviderConfigurationInput,
) (*usecase.OpenIDProviderConfigurationOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	endpoint := "http://localhost:8080"

	issuer, _ := url.Parse(endpoint)

	authorization, _ := url.Parse(fmt.Sprintf("%s/op/authorize", endpoint))

	token, _ := url.Parse(fmt.Sprintf("%s/op/token", endpoint))

	userinfo, _ := url.Parse(fmt.Sprintf("%s/op/userinfo", endpoint))

	jwks, _ := url.Parse(fmt.Sprintf("%s/op/certs", endpoint))

	revocation, _ := url.Parse(fmt.Sprintf("%s/op/revoke", endpoint))

	return &usecase.OpenIDProviderConfigurationOutput{
		Issuer:                *issuer,
		AuthorizationEndpoint: *authorization,
		TokenEndpoint:         *token,
		UserinfoEndpoint:      *userinfo,
		JwksURI:               *jwks,
		RevocationEndpoint:    *revocation,
	}, nil
}

// Autorize implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Autorize(
	ctx context.Context,
	input usecase.OpenIDProviderAuthorizeInput,
) (*usecase.OpenIDProviderAuthorizeOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	cli, err := op.client.Find(ctx, input.ClientID)
	if err != nil {
		return nil, derror.ErrUnauthorizedClient
	}

	slog.Info(fmt.Sprintf("%+v", cli))

	if cli.RedirectURI != input.RedirectURI {
		slog.Warn(fmt.Sprintf("input: %v, got: %v", input.RedirectURI, cli.RedirectURI))
		return nil, derror.ErrInvalidRequest
	}

	var buf bytes.Buffer

	buf.WriteString(input.LoginURL)

	id := uuid.NewString()

	if err := model.ValidateScope(input.Scope); err != nil {
		return nil, errors.Wrap(err, "invalid_scope")
	}

	if err := op.param.Set(ctx, id, model.AuthorizeParam{
		RedirectURI: input.RedirectURI,
		State:       input.State,
		Scope:       input.Scope,
		ClientID:    input.ClientID,
		Nonce:       input.Nonce,
	}, time.Second); err != nil {
		return nil, derror.ErrServerError
	}

	values := url.Values{
		"auth_request_id": {id},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	redirect, _ := url.ParseRequestURI(buf.String())

	return &usecase.OpenIDProviderAuthorizeOutput{
		RedirectURI: *redirect,
	}, nil
}

// LoginVeiw implements usecase.OpenIDProviider.
func (*OpenIDProvider) LoginVeiw(
	ctx context.Context,
	input usecase.OpenIDProviderLoginViewInput,
) (*usecase.OpenIDProviderLoginViewOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	var loginTmpl, _ = template.New("login").Parse(model.LoginVeiw)

	data := &struct {
		ID    string
		Error string
	}{
		ID:    input.AuthRequestID,
		Error: "",
	}

	buf := new(bytes.Buffer)

	if err := loginTmpl.Execute(buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return &usecase.OpenIDProviderLoginViewOutput{
		Data: buf,
	}, nil
}

// Login implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Login(
	ctx context.Context,
	input usecase.OpenIDProviderLoginInput,
) (*usecase.OpenIDProviderLoginOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	account, err := op.account.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	hashPass, err := op.account.FindPassword(ctx, account.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find password: %w", err)
	}

	// FIXME: 微妙 ...
	account.HashPass = hashPass

	if err != account.ComparePassword(input.Password) {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	if err := op.loggedin.Set(ctx, input.ID, model.LoggedIn{
		AccountID: account.ID,
	}, time.Minute); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	var buf bytes.Buffer

	buf.WriteString(input.CallbackURL)

	values := url.Values{
		"id": {input.ID},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	redirect, _ := url.ParseRequestURI(buf.String())

	return &usecase.OpenIDProviderLoginOutput{
		RedirectURI: *redirect,
	}, nil
}

// Callback implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Callback(
	ctx context.Context,
	input usecase.OpenIDProviderCallbackInput,
) (*usecase.OpenIDProviderCallbackOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	var buf bytes.Buffer

	val, err := op.param.Get(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get param cache: %w", err)
	}

	if _, err := op.loggedin.Get(ctx, input.ID); err != nil {
		return nil, fmt.Errorf("failed to get logged in cache: %w", err)
	}

	buf.WriteString(val.RedirectURI)

	values := url.Values{
		"code":  {input.ID},
		"state": {val.State},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	redirect, _ := url.ParseRequestURI(buf.String())

	return &usecase.OpenIDProviderCallbackOutput{
		RedirectURI: *redirect,
	}, nil
}

// AuthorizationCodeGrant implements usecase.OpenIDProviider.
func (op *OpenIDProvider) AuthorizationCodeGrant(
	ctx context.Context,
	input usecase.OpenIDProviderAuthorizationCodeGrantInput,
) (*usecase.OpenIDProviderAuthorizationCodeGrantOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	cli, err := op.client.Find(ctx, input.ClientID)
	if err != nil {
		return nil, derror.ErrUnauthorizedClient
	}

	hashSec, err := op.client.FindSecret(ctx, input.ClientID)
	if err != nil {
		return nil, fmt.Errorf("failed to find secret: %w", err)
	}

	// FIXME: 微妙 ...
	cli.HashSec = hashSec

	if err := cli.CompareSecret(input.ClientSecret); err != nil {
		return nil, fmt.Errorf("failed to compare secret: %w", err)
	}

	loggedin, err := op.loggedin.Get(ctx, input.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to get logged in cache: %w", err)
	}

	param, err := op.param.Get(ctx, input.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to get param cache: %w", err)
	}

	account, err := op.account.Find(ctx, loggedin.AccountID)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	at := model.GenerateAccessToken(
		input.Issuer,
		loggedin.AccountID,
		param.ClientID,
		"jti",
		param.Scope,
		param.ClientID,
	).JWT(input.AccessTokenSign)

	if err := op.accessToken.Set(ctx, at, struct{}{}, time.Hour); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	rt := model.GenerateRefreshToken().Base64()

	if err := op.refreshToken.Set(ctx, rt, account.ID, 24*time.Hour); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	var (
		profile *string
		email   *string
	)

	if slices.Contains(param.Scope, "profile") {
		profile = &account.Name
	}

	if slices.Contains(param.Scope, "email") {
		email = &account.Email
	}

	it := model.GenerateIDToken(
		input.Issuer,
		loggedin.AccountID,
		param.ClientID,
		param.Nonce,
		profile,
		email,
	).RSA256(input.IDTokenSignKey)

	return &usecase.OpenIDProviderAuthorizationCodeGrantOutput{
		TokenType:    "Bearer",
		AccessToken:  at,
		RefreshToken: rt,
		IDToken:      it,
		ExpiresIn:    3600,
		Scope:        param.Scope,
	}, nil
}

// RefreshTkenGrant implements usecase.OpenIDProviider.
func (op *OpenIDProvider) RefreshTkenGrant(
	ctx context.Context,
	input usecase.OpenIDProviderRefreshTokenGrantInput,
) (*usecase.OpenIDProviderRefreshTokenGrantOutput, error) {
	id, err := op.refreshToken.Get(ctx, input.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get cache: %w", err)
	}

	account, err := op.account.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	at := model.GenerateAccessToken(
		input.Issuer,
		account.ID,
		input.ClientID,
		"jti",
		input.Scope,
		input.ClientID,
	).JWT(input.AccessTokenSign)

	if err := op.accessToken.Set(ctx, at, struct{}{}, time.Hour); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	rt := model.GenerateRefreshToken().Base64()

	if err := op.refreshToken.Set(ctx, rt, account.ID, 24*time.Hour); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	var (
		profile *string
		email   *string
	)

	if slices.Contains(input.Scope, "profile") {
		profile = &account.Name
	}

	if slices.Contains(input.Scope, "email") {
		email = &account.Email
	}

	it := model.GenerateIDToken(
		input.Issuer,
		account.ID,
		input.ClientID,
		"",
		profile,
		email,
	).RSA256(input.IDTokenSignKey)

	return &usecase.OpenIDProviderRefreshTokenGrantOutput{
		TokenType:    "Bearer",
		AccessToken:  at,
		RefreshToken: rt,
		IDToken:      it,
		ExpiresIn:    3600,
		Scope:        input.Scope,
	}, nil
}

// Userinfo implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Userinfo(
	ctx context.Context,
	input usecase.OpenIDProviderUserinfoInput,
) (*usecase.OpenIDProviderUserinfoOutput, error) {
	account, err := op.account.Find(ctx, input.AccessToken.Sub)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	output := &usecase.OpenIDProviderUserinfoOutput{
		Sub: account.ID,
	}

	if slices.Contains(input.AccessToken.Scope, "profile") {
		output.Profile = &account.Name
	}

	if slices.Contains(input.AccessToken.Scope, "email") {
		output.Email = &account.Email
	}

	return output, nil
}

// Certs implements usecase.OpenIDProviider.
func (*OpenIDProvider) Certs(
	ctx context.Context,
	input usecase.OpenIDProviderCertsInput,
) (*usecase.OpenIDProviderCertsOutput, error) {
	data := make([]byte, 8)

	binary.BigEndian.PutUint64(data, uint64(input.PublicKey.E))

	i := 0
	for ; i < len(data); i++ {
		if data[i] != 0x0 {
			break
		}
	}

	e := base64.RawURLEncoding.EncodeToString(data[i:])

	return &usecase.OpenIDProviderCertsOutput{
		Kid: "12345678",
		Kty: "RSA",
		Use: "sig",
		Alg: "RS256",
		N:   base64.RawURLEncoding.EncodeToString(input.PublicKey.N.Bytes()),
		E:   e,
	}, nil
}

// Revoke implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Revoke(
	ctx context.Context,
	input usecase.OpenIDProviderRevokeInput,
) (*usecase.OpenIDProviderRevokeOutput, error) {
	switch input.Hint {
	case "access_token":
		if err := op.accessToken.Del(ctx, input.Token); err != nil {
			return nil, fmt.Errorf("failed to del cache: %w", err)
		}
	case "refresh_token":
		if err := op.refreshToken.Del(ctx, input.Token); err != nil {
			return nil, fmt.Errorf("failed to del cache: %w", err)
		}
	default:
		_ = op.refreshToken.Del(ctx, input.Token)
		_ = op.accessToken.Del(ctx, input.Token)
	}

	return &usecase.OpenIDProviderRevokeOutput{}, nil
}
