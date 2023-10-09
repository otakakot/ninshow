package interactor

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
	"github.com/otakakot/ninshow/pkg/log"
	"golang.org/x/exp/slices"
)

var _ usecase.OpenIDProviider = (*OpenIDProvider)(nil)

type OpenIDProvider struct {
	account      repository.Account
	param        repository.Cache[model.AuthorizeParam]
	loggedin     repository.Cache[model.LoggedIn]
	accessToken  repository.Cache[struct{}]
	refreshToken repository.Cache[struct{}]
}

func NewOpenIDProvider(
	account repository.Account,
	param repository.Cache[model.AuthorizeParam],
	loggedin repository.Cache[model.LoggedIn],
	accessToken repository.Cache[struct{}],
	refreshToken repository.Cache[struct{}],
) *OpenIDProvider {
	return &OpenIDProvider{
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

	jwks, _ := url.Parse(fmt.Sprintf("%s/op/jwks", endpoint))

	revocation, _ := url.Parse(fmt.Sprintf("%s/op/revoke", endpoint))

	return &usecase.OpenIDProviderConfigurationOutput{
		Issuer:                *issuer,
		AuthorizationEndpoint: *authorization,
		TokenEndpoint:         *token,
		UserinfoEndpoint:      *userinfo,
		JwksURL:               *jwks,
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

	var buf bytes.Buffer

	buf.WriteString(input.LoginURL)

	id := uuid.NewString()

	if err := op.param.Set(ctx, id, model.AuthorizeParam{
		RedirectURI: input.RedirectURI,
		State:       input.State,
		Scope:       input.Scope,
		ClientID:    input.ClientID,
		Nonce:       input.Nonce,
	}, time.Second); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
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

	const tmp = `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Login</title>
		</head>
		<body bgcolor="black" style="display: flex; align-items: center; justify-content: center; height: 100vh;">
			<form method="POST" action="/op/login" style="height: 200px; width: 200px;">

				<input type="hidden" name="id" value="{{.ID}}">

				<div style="color:white;">
					<label for="username">Username:</label>
					<input id="username" name="username" style="width: 100%">
				</div>

				<div style="color:white;">
					<label for="password">Password:</label>
					<input id="password" name="password" style="width: 100%" type="password">
				</div>

				<p style="color:red; min-height: 1rem;">{{.Error}}</p>

				<button type="submit">Login</button>
			</form>
		</body>
	</html>
	`

	var loginTmpl, _ = template.New("login").Parse(tmp)

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

	account, err := op.account.FindByUsername(ctx, input.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

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

// Token implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Token(
	ctx context.Context,
	input usecase.OpenIDProviderTokenInput,
) (*usecase.OpenIDProviderTokenOutput, error) {
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

	if err := op.refreshToken.Set(ctx, rt, struct{}{}, 24*time.Hour); err != nil {
		return nil, fmt.Errorf("failed to set cache: %w", err)
	}

	var (
		profile *string
		email   *string
	)

	if slices.Contains(param.Scope, "profile") {
		profile = &account.Username
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

	return &usecase.OpenIDProviderTokenOutput{
		TokenType:    "Bearer",
		AccessToken:  at,
		RefreshToken: rt,
		IDToken:      it,
		ExpiresIn:    3600,
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
		output.Profile = &account.Username
	}

	if slices.Contains(input.AccessToken.Scope, "email") {
		output.Email = &account.Email
	}

	return output, nil
}
