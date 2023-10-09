package usecase

import (
	"context"
	"crypto/rsa"
	"io"
	"net/url"
)

type OpenIDProviider interface {
	Configuration(context.Context, OpenIDProviderConfigurationInput) (*OpenIDProviderConfigurationOutput, error)
	Autorize(context.Context, OpenIDProviderAuthorizeInput) (*OpenIDProviderAuthorizeOutput, error)
	LoginVeiw(context.Context, OpenIDProviderLoginViewInput) (*OpenIDProviderLoginViewOutput, error)
	Login(context.Context, OpenIDProviderLoginInput) (*OpenIDProviderLoginOutput, error)
	Callback(context.Context, OpenIDProviderCallbackInput) (*OpenIDProviderCallbackOutput, error)
	Token(context.Context, OpenIDProviderTokenInput) (*OpenIDProviderTokenOutput, error)
	Userinfo(context.Context, OpenIDProviderUserinfoInput) (*OpenIDProviderUserinfoOutput, error)
}

type OpenIDProviderConfigurationInput struct{}

type OpenIDProviderConfigurationOutput struct {
	Issuer                url.URL
	AuthorizationEndpoint url.URL
	TokenEndpoint         url.URL
	UserinfoEndpoint      url.URL
	JwksURL               url.URL
	RevocationEndpoint    url.URL
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

type OpenIDProviderLoginViewInput struct {
	AuthRequestID string
}

type OpenIDProviderLoginViewOutput struct {
	Data io.Reader
}

type OpenIDProviderLoginInput struct {
	ID          string
	Username    string
	Password    string
	CallbackURL string
}

type OpenIDProviderLoginOutput struct {
	RedirectURI url.URL
}

type OpenIDProviderCallbackInput struct {
	ID string
}

type OpenIDProviderCallbackOutput struct {
	RedirectURI url.URL
}

type OpenIDProviderTokenInput struct {
	Issuer          string
	Code            string
	AccessTokenSign string
	IDTokenSignKey  *rsa.PrivateKey
}

type OpenIDProviderTokenOutput struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
	IDToken      string
	ExpiresIn    int
}

type OpenIDProviderUserinfoInput struct{}

type OpenIDProviderUserinfoOutput struct{}
