package usecase

import (
	"context"
	"crypto/rsa"
	"io"
	"net/url"

	"github.com/otakakot/ninshow/internal/domain/model"
)

type OpenIDProviider interface {
	Configuration(context.Context, OpenIDProviderConfigurationInput) (*OpenIDProviderConfigurationOutput, error)
	Autorize(context.Context, OpenIDProviderAuthorizeInput) (*OpenIDProviderAuthorizeOutput, error)
	LoginVeiw(context.Context, OpenIDProviderLoginViewInput) (*OpenIDProviderLoginViewOutput, error)
	Login(context.Context, OpenIDProviderLoginInput) (*OpenIDProviderLoginOutput, error)
	Callback(context.Context, OpenIDProviderCallbackInput) (*OpenIDProviderCallbackOutput, error)
	AuthorizationCodeGrant(context.Context, OpenIDProviderAuthorizationCodeGrantInput) (*OpenIDProviderAuthorizationCodeGrantOutput, error)
	RefreshTkenGrant(context.Context, OpenIDProviderRefreshTokenGrantInput) (*OpenIDProviderRefreshTokenGrantOutput, error)
	Userinfo(context.Context, OpenIDProviderUserinfoInput) (*OpenIDProviderUserinfoOutput, error)
	Certs(context.Context, OpenIDProviderCertsInput) (*OpenIDProviderCertsOutput, error)
	Revoke(context.Context, OpenIDProviderRevokeInput) (*OpenIDProviderRevokeOutput, error)
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

type OpenIDProviderAuthorizationCodeGrantInput struct {
	Issuer          string
	Code            string
	AccessTokenSign string
	IDTokenSignKey  *rsa.PrivateKey
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
	IDTokenSignKey  *rsa.PrivateKey
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
	PublicKey rsa.PublicKey
}

type OpenIDProviderCertsOutput struct {
	Kid string // Kid 鍵識別子
	Kty string // Kty RSAやEC等の暗号アルゴリズファミリー
	Use string // Use 公開鍵の用途
	Alg string // Alg 署名検証アルゴリズム
	N   string // N modulus 公開鍵を復元するための公開鍵の絶対値
	E   string // E exponent 公開鍵を復元するための指数値
}

type OpenIDProviderRevokeInput struct {
	Hint  string
	Token string
}

type OpenIDProviderRevokeOutput struct{}
