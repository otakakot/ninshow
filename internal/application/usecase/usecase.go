package usecase

import (
	"context"
	"net/http"
	"net/url"
)

// Identity Provider

type IdentityProvider interface {
	Signup(context.Context, IdentityProviderSignupInput) (*IdentityProviderSignupOutput, error)
	Signin(context.Context, IdentityProviderSigninInput) (*IdentityProviderSigninOutput, error)
}

type IdentityProviderSignupInput struct {
	Email    string
	Username string
	Password string
}

type IdentityProviderSignupOutput struct{}

type IdentityProviderSigninInput struct {
	Username string
	Password string
}

type IdentityProviderSigninOutput struct{}

// OpenID Provider

type OpenIDProviider interface {
	Autorize(context.Context, OpenIDProviderAuthorizeInput) (*OpenIDProviderAuthorizeOutput, error)
}

type OpenIDProviderAuthorizeInput struct {
	LoginURL     string
	ResponseType string
	Scope        []string
	ClientID     string
	RedirectURI  string
	State        string
	Nonce        string
}

type OpenIDProviderAuthorizeOutput struct {
	RedirectURI url.URL
}

// Relying Party

type RelyingParty interface {
	Login(context.Context, RelyingPartyLoginInput) (*RelyingPartyLoginOutput, error)
}

type RelyingPartyLoginInput struct {
	OIDCEndpoint string   // OP の認証エンドポイント
	ClientID     string   // OP に登録してあるクライアントID
	RedirectURI  string   // OP に登録してある認証成功後にリダイレクトさせるURI
	Scope        []string // OP に登録してあるスコープ
}

type RelyingPartyLoginOutput struct {
	Cookie      *http.Cookie
	RedirectURI url.URL
}
