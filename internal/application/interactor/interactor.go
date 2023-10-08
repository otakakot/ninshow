package interactor

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
)

// Identity Provider

var _ usecase.IdentityProvider = (*IdentityProvider)(nil)

type IdentityProvider struct {
	account repository.Account
}

func NewIdentityProvider(
	acccount repository.Account,
) *IdentityProvider {
	return &IdentityProvider{
		account: acccount,
	}
}

// Signup implements usecase.IdentityProvider.
func (idp *IdentityProvider) Signup(
	ctx context.Context,
	input usecase.IdentityProviderSignupInput,
) (*usecase.IdentityProviderSignupOutput, error) {
	account, err := model.SingupAccount(
		input.Username,
		input.Email,
		input.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to signup account: %w", err)
	}

	if err := idp.account.Save(ctx, *account); err != nil {
		return nil, fmt.Errorf("failed to save account: %w", err)
	}

	return &usecase.IdentityProviderSignupOutput{}, nil
}

// Signin implements usecase.IdentityProvider.
func (idp *IdentityProvider) Signin(
	ctx context.Context,
	input usecase.IdentityProviderSigninInput,
) (*usecase.IdentityProviderSigninOutput, error) {
	account, err := idp.account.Find(ctx, input.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	if err != account.ComparePassword(input.Password) {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	return &usecase.IdentityProviderSigninOutput{}, nil
}

// OpenID Provider

var _ usecase.OpenIDProviider = (*OpenIDProvider)(nil)

type OpenIDProvider struct {
	kvs repository.Cache[any]
}

func NewOpenIDProvider(
	kvs repository.Cache[any],
) *OpenIDProvider {
	return &OpenIDProvider{
		kvs: kvs,
	}
}

// Autorize implements usecase.OpenIDProviider.
func (op *OpenIDProvider) Autorize(
	ctx context.Context,
	input usecase.OpenIDProviderAuthorizeInput,
) (*usecase.OpenIDProviderAuthorizeOutput, error) {
	var buf bytes.Buffer

	buf.WriteString(input.LoginURL)

	id := uuid.NewString()

	if err := op.kvs.Set(ctx, id, input, time.Second); err != nil {
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

// Relying Party

var _ usecase.RelyingParty = (*RelyingParty)(nil)

type RelyingParty struct{}

func NewRelyingParty() *RelyingParty {
	return &RelyingParty{}
}

// Login implements usecase.RelyingParty.
func (*RelyingParty) Login(
	_ context.Context,
	input usecase.RelyingPartyLoginInput,
) (*usecase.RelyingPartyLoginOutput, error) {
	var buf bytes.Buffer

	buf.WriteString(input.OIDCEndpoint)

	state := uuid.NewString()

	cookie := &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	values := url.Values{
		"response_type": {"code"},                         // Authorization Flow なので code を指定
		"client_id":     {input.ClientID},                 // RPを識別するためのID OPに登録しておく必要がある
		"redirect_uri":  {input.RedirectURI},              // ログイン後にリダイレクトさせるURL OPに登録しておく必要がある
		"scope":         {strings.Join(input.Scope, " ")}, // RPが要求するスコープ OPに登録しておく必要がある
		"state":         {state},                          // CSRF対策のためのstate
		"nonce":         {uuid.NewString()},               // CSRF対策のためのnonce
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	redirect, _ := url.ParseRequestURI(buf.String())

	return &usecase.RelyingPartyLoginOutput{
		Cookie:      cookie,
		RedirectURI: *redirect,
	}, nil
}
