package usecase

import (
	"context"
	"crypto/rsa"
	"io"
	"net/url"

	"github.com/otakakot/ninshow/internal/domain/model"
)

type OpenIDProviider interface {
	Configuration(ctx context.Context, input OpenIDProviderConfigurationInput) (*OpenIDProviderConfigurationOutput, error)
	Authorize(ctx context.Context, input OpenIDProviderAuthorizeInput) (*OpenIDProviderAuthorizeOutput, error)
	LoginVeiw(ctx context.Context, input OpenIDProviderLoginViewInput) (*OpenIDProviderLoginViewOutput, error)
	Login(ctx context.Context, input OpenIDProviderLoginInput) (*OpenIDProviderLoginOutput, error)
	Callback(ctx context.Context, input OpenIDProviderCallbackInput) (*OpenIDProviderCallbackOutput, error)
	AuthorizationCodeGrant(ctx context.Context, input OpenIDProviderAuthorizationCodeGrantInput) (*OpenIDProviderAuthorizationCodeGrantOutput, error)
	RefreshTokenGrant(ctx context.Context, input OpenIDProviderRefreshTokenGrantInput) (*OpenIDProviderRefreshTokenGrantOutput, error)
	Userinfo(ctx context.Context, input OpenIDProviderUserinfoInput) (*OpenIDProviderUserinfoOutput, error)
	Certs(ctx context.Context, input OpenIDProviderCertsInput) (*OpenIDProviderCertsOutput, error)
	Revoke(ctx context.Context, input OpenIDProviderRevokeInput) (*OpenIDProviderRevokeOutput, error)
}

type OpenIDProviderConfigurationInput struct{}

type OpenIDProviderConfigurationOutput struct {
	Issuer                url.URL
	AuthorizationEndpoint url.URL
	TokenEndpoint         url.URL
	UserinfoEndpoint      url.URL
	JwksURI               url.URL
	RevocationEndpoint    url.URL
}

type OpenIDProviderAuthorizeInput struct {
	LoginURL      string
	ResponseType  string
	Scope         []string
	ClientID      string
	RedirectURI   string
	State         string
	Nonce         string
	CodeChallenge *model.CodeChallenge
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
	Email       string
	Password    string
	CallbackURL string
}

type OpenIDProviderLoginOutput struct {
	RedirectURI url.URL
}

type OpenIDProviderCallbackInput struct {
	ID             string
	Issuer         string
	IDTokenSignKey *rsa.PrivateKey
}

type OpenIDProviderCallbackOutput struct {
	RedirectURI url.URL
}

type OpenIDProviderAuthorizationCodeGrantInput struct {
	ClientID        string
	ClientSecret    string
	Issuer          string
	Code            string
	CodeVerifier    string
	AccessTokenSign string
}

type OpenIDProviderAuthorizationCodeGrantOutput struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
	IDToken      string
	ExpiresIn    int
	Scope        []string
}

type OpenIDProviderRefreshTokenGrantInput struct {
	RefreshToken    string
	ClientID        string
	Scope           []string
	Issuer          string
	AccessTokenSign string
}

type OpenIDProviderRefreshTokenGrantOutput struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
	IDToken      string
	ExpiresIn    int
	Scope        []string
}

type OpenIDProviderUserinfoInput struct {
	AccessToken model.AccessToken
	Scope       []string
}

type OpenIDProviderUserinfoOutput struct {
	Sub     string
	Profile *string
	Email   *string
}

type OpenIDProviderCertsInput struct {
}

type OpenIDProviderCertsOutput struct {
	Certs []model.Cert
}

type OpenIDProviderRevokeInput struct {
	Hint  string
	Token string
}

type OpenIDProviderRevokeOutput struct{}
