// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/url"

	"github.com/go-faster/errors"
)

type Bearer struct {
	Token string
}

// GetToken returns the value of Token.
func (s *Bearer) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Bearer) SetToken(val string) {
	s.Token = val
}

// HealthInternalServerError is response for Health operation.
type HealthInternalServerError struct{}

func (*HealthInternalServerError) healthRes() {}

// HealthOK is response for Health operation.
type HealthOK struct{}

func (*HealthOK) healthRes() {}

// Ref: #/components/schemas/IdPSigninRequestSchema
type IdPSigninRequestSchema struct {
	// Username.
	Username string `json:"username"`
	// Password.
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *IdPSigninRequestSchema) GetUsername() string {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *IdPSigninRequestSchema) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *IdPSigninRequestSchema) SetUsername(val string) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *IdPSigninRequestSchema) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/IdPSignupRequestSchema
type IdPSignupRequestSchema struct {
	// Username.
	Username string `json:"username"`
	Email    string `json:"email"`
	// Password.
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *IdPSignupRequestSchema) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *IdPSignupRequestSchema) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *IdPSignupRequestSchema) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *IdPSignupRequestSchema) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *IdPSignupRequestSchema) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *IdPSignupRequestSchema) SetPassword(val string) {
	s.Password = val
}

// IdpSigninInternalServerError is response for IdpSignin operation.
type IdpSigninInternalServerError struct{}

func (*IdpSigninInternalServerError) idpSigninRes() {}

// IdpSigninOK is response for IdpSignin operation.
type IdpSigninOK struct{}

func (*IdpSigninOK) idpSigninRes() {}

// IdpSigninUnauthorized is response for IdpSignin operation.
type IdpSigninUnauthorized struct{}

func (*IdpSigninUnauthorized) idpSigninRes() {}

// IdpSignupInternalServerError is response for IdpSignup operation.
type IdpSignupInternalServerError struct{}

func (*IdpSignupInternalServerError) idpSignupRes() {}

// IdpSignupOK is response for IdpSignup operation.
type IdpSignupOK struct{}

func (*IdpSignupOK) idpSignupRes() {}

// Jwk set key.
// Ref: #/components/schemas/OPJWKSetKey
type OPJWKSetKey struct {
	// 鍵識別子.
	Kid string `json:"kid"`
	// RSAやEC等の暗号アルゴリズファミリー.
	Kty string `json:"kty"`
	// 公開鍵の用途.
	Use string `json:"use"`
	// 署名検証アルゴリズム.
	Alg string `json:"alg"`
	// Modulus 公開鍵を復元するための公開鍵の絶対値.
	N string `json:"n"`
	// Exponent 公開鍵を復元するための指数値.
	E string `json:"e"`
}

// GetKid returns the value of Kid.
func (s *OPJWKSetKey) GetKid() string {
	return s.Kid
}

// GetKty returns the value of Kty.
func (s *OPJWKSetKey) GetKty() string {
	return s.Kty
}

// GetUse returns the value of Use.
func (s *OPJWKSetKey) GetUse() string {
	return s.Use
}

// GetAlg returns the value of Alg.
func (s *OPJWKSetKey) GetAlg() string {
	return s.Alg
}

// GetN returns the value of N.
func (s *OPJWKSetKey) GetN() string {
	return s.N
}

// GetE returns the value of E.
func (s *OPJWKSetKey) GetE() string {
	return s.E
}

// SetKid sets the value of Kid.
func (s *OPJWKSetKey) SetKid(val string) {
	s.Kid = val
}

// SetKty sets the value of Kty.
func (s *OPJWKSetKey) SetKty(val string) {
	s.Kty = val
}

// SetUse sets the value of Use.
func (s *OPJWKSetKey) SetUse(val string) {
	s.Use = val
}

// SetAlg sets the value of Alg.
func (s *OPJWKSetKey) SetAlg(val string) {
	s.Alg = val
}

// SetN sets the value of N.
func (s *OPJWKSetKey) SetN(val string) {
	s.N = val
}

// SetE sets the value of E.
func (s *OPJWKSetKey) SetE(val string) {
	s.E = val
}

// Https://openid-foundation-japan.github.io/rfc7517.ja.html#anchor5.
// Ref: #/components/schemas/OPJWKSetResponseSchema
type OPJWKSetResponseSchema struct {
	Keys []OPJWKSetKey `json:"keys"`
}

// GetKeys returns the value of Keys.
func (s *OPJWKSetResponseSchema) GetKeys() []OPJWKSetKey {
	return s.Keys
}

// SetKeys sets the value of Keys.
func (s *OPJWKSetResponseSchema) SetKeys(val []OPJWKSetKey) {
	s.Keys = val
}

func (*OPJWKSetResponseSchema) opCertsRes() {}

// Ref: #/components/schemas/OPOpenIDConfigurationResponseSchema
type OPOpenIDConfigurationResponseSchema struct {
	// Http://localhost:8080/op.
	Issuer url.URL `json:"issuer"`
	// Http://localhost:8080/op/authorize.
	AuthorizationEndpoint url.URL `json:"authorization_endpoint"`
	// Http://localhost:8080/op/token.
	TokenEndpoint url.URL `json:"token_endpoint"`
	// Http://localhost:8080/op/userinfo.
	UserinfoEndpoint url.URL `json:"userinfo_endpoint"`
	// Http://localhost:8080/op/certs.
	JwksURL url.URL `json:"jwks_url"`
	// Http://localhost:8080/op/revoke.
	RevocationEndpoint url.URL `json:"revocation_endpoint"`
}

// GetIssuer returns the value of Issuer.
func (s *OPOpenIDConfigurationResponseSchema) GetIssuer() url.URL {
	return s.Issuer
}

// GetAuthorizationEndpoint returns the value of AuthorizationEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) GetAuthorizationEndpoint() url.URL {
	return s.AuthorizationEndpoint
}

// GetTokenEndpoint returns the value of TokenEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) GetTokenEndpoint() url.URL {
	return s.TokenEndpoint
}

// GetUserinfoEndpoint returns the value of UserinfoEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) GetUserinfoEndpoint() url.URL {
	return s.UserinfoEndpoint
}

// GetJwksURL returns the value of JwksURL.
func (s *OPOpenIDConfigurationResponseSchema) GetJwksURL() url.URL {
	return s.JwksURL
}

// GetRevocationEndpoint returns the value of RevocationEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) GetRevocationEndpoint() url.URL {
	return s.RevocationEndpoint
}

// SetIssuer sets the value of Issuer.
func (s *OPOpenIDConfigurationResponseSchema) SetIssuer(val url.URL) {
	s.Issuer = val
}

// SetAuthorizationEndpoint sets the value of AuthorizationEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) SetAuthorizationEndpoint(val url.URL) {
	s.AuthorizationEndpoint = val
}

// SetTokenEndpoint sets the value of TokenEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) SetTokenEndpoint(val url.URL) {
	s.TokenEndpoint = val
}

// SetUserinfoEndpoint sets the value of UserinfoEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) SetUserinfoEndpoint(val url.URL) {
	s.UserinfoEndpoint = val
}

// SetJwksURL sets the value of JwksURL.
func (s *OPOpenIDConfigurationResponseSchema) SetJwksURL(val url.URL) {
	s.JwksURL = val
}

// SetRevocationEndpoint sets the value of RevocationEndpoint.
func (s *OPOpenIDConfigurationResponseSchema) SetRevocationEndpoint(val url.URL) {
	s.RevocationEndpoint = val
}

func (*OPOpenIDConfigurationResponseSchema) opOpenIDConfigurationRes() {}

// Https://openid-foundation-japan.github.io/rfc7009.ja.html#anchor2.
// Ref: #/components/schemas/OPRevokeRequestSchema
type OPRevokeRequestSchema struct {
	// Token.
	Token string `json:"token"`
	// Token_type_hint.
	TokenTypeHint OptOPRevokeRequestSchemaTokenTypeHint `json:"token_type_hint"`
}

// GetToken returns the value of Token.
func (s *OPRevokeRequestSchema) GetToken() string {
	return s.Token
}

// GetTokenTypeHint returns the value of TokenTypeHint.
func (s *OPRevokeRequestSchema) GetTokenTypeHint() OptOPRevokeRequestSchemaTokenTypeHint {
	return s.TokenTypeHint
}

// SetToken sets the value of Token.
func (s *OPRevokeRequestSchema) SetToken(val string) {
	s.Token = val
}

// SetTokenTypeHint sets the value of TokenTypeHint.
func (s *OPRevokeRequestSchema) SetTokenTypeHint(val OptOPRevokeRequestSchemaTokenTypeHint) {
	s.TokenTypeHint = val
}

// Token_type_hint.
type OPRevokeRequestSchemaTokenTypeHint string

const (
	OPRevokeRequestSchemaTokenTypeHintAccessToken  OPRevokeRequestSchemaTokenTypeHint = "access_token"
	OPRevokeRequestSchemaTokenTypeHintRefreshToken OPRevokeRequestSchemaTokenTypeHint = "refresh_token"
)

// AllValues returns all OPRevokeRequestSchemaTokenTypeHint values.
func (OPRevokeRequestSchemaTokenTypeHint) AllValues() []OPRevokeRequestSchemaTokenTypeHint {
	return []OPRevokeRequestSchemaTokenTypeHint{
		OPRevokeRequestSchemaTokenTypeHintAccessToken,
		OPRevokeRequestSchemaTokenTypeHintRefreshToken,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s OPRevokeRequestSchemaTokenTypeHint) MarshalText() ([]byte, error) {
	switch s {
	case OPRevokeRequestSchemaTokenTypeHintAccessToken:
		return []byte(s), nil
	case OPRevokeRequestSchemaTokenTypeHintRefreshToken:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OPRevokeRequestSchemaTokenTypeHint) UnmarshalText(data []byte) error {
	switch OPRevokeRequestSchemaTokenTypeHint(data) {
	case OPRevokeRequestSchemaTokenTypeHintAccessToken:
		*s = OPRevokeRequestSchemaTokenTypeHintAccessToken
		return nil
	case OPRevokeRequestSchemaTokenTypeHintRefreshToken:
		*s = OPRevokeRequestSchemaTokenTypeHintRefreshToken
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/OPTokenRequestSchema
type OPTokenRequestSchema struct {
	// Grant_type.
	GrantType OPTokenRequestSchemaGrantType `json:"grant_type"`
	// Code.
	Code string `json:"code"`
	// Http://localhost:8080/rp/callback.
	RedirectURI url.URL `json:"redirect_uri"`
}

// GetGrantType returns the value of GrantType.
func (s *OPTokenRequestSchema) GetGrantType() OPTokenRequestSchemaGrantType {
	return s.GrantType
}

// GetCode returns the value of Code.
func (s *OPTokenRequestSchema) GetCode() string {
	return s.Code
}

// GetRedirectURI returns the value of RedirectURI.
func (s *OPTokenRequestSchema) GetRedirectURI() url.URL {
	return s.RedirectURI
}

// SetGrantType sets the value of GrantType.
func (s *OPTokenRequestSchema) SetGrantType(val OPTokenRequestSchemaGrantType) {
	s.GrantType = val
}

// SetCode sets the value of Code.
func (s *OPTokenRequestSchema) SetCode(val string) {
	s.Code = val
}

// SetRedirectURI sets the value of RedirectURI.
func (s *OPTokenRequestSchema) SetRedirectURI(val url.URL) {
	s.RedirectURI = val
}

// Grant_type.
type OPTokenRequestSchemaGrantType string

const (
	OPTokenRequestSchemaGrantTypeAuthorizationCode                     OPTokenRequestSchemaGrantType = "authorization_code"
	OPTokenRequestSchemaGrantTypeRefreshToken                          OPTokenRequestSchemaGrantType = "refresh_token"
	OPTokenRequestSchemaGrantTypeClientCredentials                     OPTokenRequestSchemaGrantType = "client_credentials"
	OPTokenRequestSchemaGrantTypePassword                              OPTokenRequestSchemaGrantType = "password"
	OPTokenRequestSchemaGrantTypeUrnIetfParamsOAuthGrantTypeDeviceCode OPTokenRequestSchemaGrantType = "urn:ietf:params:oauth:grant-type:device_code"
)

// AllValues returns all OPTokenRequestSchemaGrantType values.
func (OPTokenRequestSchemaGrantType) AllValues() []OPTokenRequestSchemaGrantType {
	return []OPTokenRequestSchemaGrantType{
		OPTokenRequestSchemaGrantTypeAuthorizationCode,
		OPTokenRequestSchemaGrantTypeRefreshToken,
		OPTokenRequestSchemaGrantTypeClientCredentials,
		OPTokenRequestSchemaGrantTypePassword,
		OPTokenRequestSchemaGrantTypeUrnIetfParamsOAuthGrantTypeDeviceCode,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s OPTokenRequestSchemaGrantType) MarshalText() ([]byte, error) {
	switch s {
	case OPTokenRequestSchemaGrantTypeAuthorizationCode:
		return []byte(s), nil
	case OPTokenRequestSchemaGrantTypeRefreshToken:
		return []byte(s), nil
	case OPTokenRequestSchemaGrantTypeClientCredentials:
		return []byte(s), nil
	case OPTokenRequestSchemaGrantTypePassword:
		return []byte(s), nil
	case OPTokenRequestSchemaGrantTypeUrnIetfParamsOAuthGrantTypeDeviceCode:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OPTokenRequestSchemaGrantType) UnmarshalText(data []byte) error {
	switch OPTokenRequestSchemaGrantType(data) {
	case OPTokenRequestSchemaGrantTypeAuthorizationCode:
		*s = OPTokenRequestSchemaGrantTypeAuthorizationCode
		return nil
	case OPTokenRequestSchemaGrantTypeRefreshToken:
		*s = OPTokenRequestSchemaGrantTypeRefreshToken
		return nil
	case OPTokenRequestSchemaGrantTypeClientCredentials:
		*s = OPTokenRequestSchemaGrantTypeClientCredentials
		return nil
	case OPTokenRequestSchemaGrantTypePassword:
		*s = OPTokenRequestSchemaGrantTypePassword
		return nil
	case OPTokenRequestSchemaGrantTypeUrnIetfParamsOAuthGrantTypeDeviceCode:
		*s = OPTokenRequestSchemaGrantTypeUrnIetfParamsOAuthGrantTypeDeviceCode
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenResponse.
// Ref: #/components/schemas/OPTokenResponseSchema
type OPTokenResponseSchema struct {
	// Access_token.
	AccessToken string `json:"access_token"`
	// Token_type.
	TokenType string `json:"token_type"`
	// Refresh_token.
	RefreshToken string `json:"refresh_token"`
	// Expires_in.
	ExpiresIn int `json:"expires_in"`
	// Id_token.
	IDToken string `json:"id_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *OPTokenResponseSchema) GetAccessToken() string {
	return s.AccessToken
}

// GetTokenType returns the value of TokenType.
func (s *OPTokenResponseSchema) GetTokenType() string {
	return s.TokenType
}

// GetRefreshToken returns the value of RefreshToken.
func (s *OPTokenResponseSchema) GetRefreshToken() string {
	return s.RefreshToken
}

// GetExpiresIn returns the value of ExpiresIn.
func (s *OPTokenResponseSchema) GetExpiresIn() int {
	return s.ExpiresIn
}

// GetIDToken returns the value of IDToken.
func (s *OPTokenResponseSchema) GetIDToken() string {
	return s.IDToken
}

// SetAccessToken sets the value of AccessToken.
func (s *OPTokenResponseSchema) SetAccessToken(val string) {
	s.AccessToken = val
}

// SetTokenType sets the value of TokenType.
func (s *OPTokenResponseSchema) SetTokenType(val string) {
	s.TokenType = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *OPTokenResponseSchema) SetRefreshToken(val string) {
	s.RefreshToken = val
}

// SetExpiresIn sets the value of ExpiresIn.
func (s *OPTokenResponseSchema) SetExpiresIn(val int) {
	s.ExpiresIn = val
}

// SetIDToken sets the value of IDToken.
func (s *OPTokenResponseSchema) SetIDToken(val string) {
	s.IDToken = val
}

func (*OPTokenResponseSchema) opTokenRes() {}

// Https://openid.net/specs/openid-connect-core-1_0.html#UserInfoResponse.
// Ref: #/components/schemas/OPUserInfoResponseSchema
type OPUserInfoResponseSchema struct {
	// Sub.
	Sub string `json:"sub"`
	// Name.
	Name string `json:"name"`
}

// GetSub returns the value of Sub.
func (s *OPUserInfoResponseSchema) GetSub() string {
	return s.Sub
}

// GetName returns the value of Name.
func (s *OPUserInfoResponseSchema) GetName() string {
	return s.Name
}

// SetSub sets the value of Sub.
func (s *OPUserInfoResponseSchema) SetSub(val string) {
	s.Sub = val
}

// SetName sets the value of Name.
func (s *OPUserInfoResponseSchema) SetName(val string) {
	s.Name = val
}

func (*OPUserInfoResponseSchema) opUserinfoRes() {}

// OpAuthorizeFound is response for OpAuthorize operation.
type OpAuthorizeFound struct {
	Location OptURI
}

// GetLocation returns the value of Location.
func (s *OpAuthorizeFound) GetLocation() OptURI {
	return s.Location
}

// SetLocation sets the value of Location.
func (s *OpAuthorizeFound) SetLocation(val OptURI) {
	s.Location = val
}

func (*OpAuthorizeFound) opAuthorizeRes() {}

// OpAuthorizeInternalServerError is response for OpAuthorize operation.
type OpAuthorizeInternalServerError struct{}

func (*OpAuthorizeInternalServerError) opAuthorizeRes() {}

type OpAuthorizeResponseType string

const (
	OpAuthorizeResponseTypeCode             OpAuthorizeResponseType = "code"
	OpAuthorizeResponseTypeIDToken          OpAuthorizeResponseType = "id_token"
	OpAuthorizeResponseTypeToken            OpAuthorizeResponseType = "token"
	OpAuthorizeResponseTypeCodeIDToken      OpAuthorizeResponseType = "code id_token"
	OpAuthorizeResponseTypeCodeToken        OpAuthorizeResponseType = "code token"
	OpAuthorizeResponseTypeIDTokenToken     OpAuthorizeResponseType = "id_token token"
	OpAuthorizeResponseTypeCodeIDTokenToken OpAuthorizeResponseType = "code id_token token"
)

// AllValues returns all OpAuthorizeResponseType values.
func (OpAuthorizeResponseType) AllValues() []OpAuthorizeResponseType {
	return []OpAuthorizeResponseType{
		OpAuthorizeResponseTypeCode,
		OpAuthorizeResponseTypeIDToken,
		OpAuthorizeResponseTypeToken,
		OpAuthorizeResponseTypeCodeIDToken,
		OpAuthorizeResponseTypeCodeToken,
		OpAuthorizeResponseTypeIDTokenToken,
		OpAuthorizeResponseTypeCodeIDTokenToken,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s OpAuthorizeResponseType) MarshalText() ([]byte, error) {
	switch s {
	case OpAuthorizeResponseTypeCode:
		return []byte(s), nil
	case OpAuthorizeResponseTypeIDToken:
		return []byte(s), nil
	case OpAuthorizeResponseTypeToken:
		return []byte(s), nil
	case OpAuthorizeResponseTypeCodeIDToken:
		return []byte(s), nil
	case OpAuthorizeResponseTypeCodeToken:
		return []byte(s), nil
	case OpAuthorizeResponseTypeIDTokenToken:
		return []byte(s), nil
	case OpAuthorizeResponseTypeCodeIDTokenToken:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OpAuthorizeResponseType) UnmarshalText(data []byte) error {
	switch OpAuthorizeResponseType(data) {
	case OpAuthorizeResponseTypeCode:
		*s = OpAuthorizeResponseTypeCode
		return nil
	case OpAuthorizeResponseTypeIDToken:
		*s = OpAuthorizeResponseTypeIDToken
		return nil
	case OpAuthorizeResponseTypeToken:
		*s = OpAuthorizeResponseTypeToken
		return nil
	case OpAuthorizeResponseTypeCodeIDToken:
		*s = OpAuthorizeResponseTypeCodeIDToken
		return nil
	case OpAuthorizeResponseTypeCodeToken:
		*s = OpAuthorizeResponseTypeCodeToken
		return nil
	case OpAuthorizeResponseTypeIDTokenToken:
		*s = OpAuthorizeResponseTypeIDTokenToken
		return nil
	case OpAuthorizeResponseTypeCodeIDTokenToken:
		*s = OpAuthorizeResponseTypeCodeIDTokenToken
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type OpAuthorizeScope string

const (
	OpAuthorizeScopeOpenid        OpAuthorizeScope = "openid"
	OpAuthorizeScopeProfile       OpAuthorizeScope = "profile"
	OpAuthorizeScopeEmail         OpAuthorizeScope = "email"
	OpAuthorizeScopeAddress       OpAuthorizeScope = "address"
	OpAuthorizeScopePhone         OpAuthorizeScope = "phone"
	OpAuthorizeScopeOfflineAccess OpAuthorizeScope = "offline_access"
)

// AllValues returns all OpAuthorizeScope values.
func (OpAuthorizeScope) AllValues() []OpAuthorizeScope {
	return []OpAuthorizeScope{
		OpAuthorizeScopeOpenid,
		OpAuthorizeScopeProfile,
		OpAuthorizeScopeEmail,
		OpAuthorizeScopeAddress,
		OpAuthorizeScopePhone,
		OpAuthorizeScopeOfflineAccess,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s OpAuthorizeScope) MarshalText() ([]byte, error) {
	switch s {
	case OpAuthorizeScopeOpenid:
		return []byte(s), nil
	case OpAuthorizeScopeProfile:
		return []byte(s), nil
	case OpAuthorizeScopeEmail:
		return []byte(s), nil
	case OpAuthorizeScopeAddress:
		return []byte(s), nil
	case OpAuthorizeScopePhone:
		return []byte(s), nil
	case OpAuthorizeScopeOfflineAccess:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OpAuthorizeScope) UnmarshalText(data []byte) error {
	switch OpAuthorizeScope(data) {
	case OpAuthorizeScopeOpenid:
		*s = OpAuthorizeScopeOpenid
		return nil
	case OpAuthorizeScopeProfile:
		*s = OpAuthorizeScopeProfile
		return nil
	case OpAuthorizeScopeEmail:
		*s = OpAuthorizeScopeEmail
		return nil
	case OpAuthorizeScopeAddress:
		*s = OpAuthorizeScopeAddress
		return nil
	case OpAuthorizeScopePhone:
		*s = OpAuthorizeScopePhone
		return nil
	case OpAuthorizeScopeOfflineAccess:
		*s = OpAuthorizeScopeOfflineAccess
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// OpCallbackFound is response for OpCallback operation.
type OpCallbackFound struct {
	Location OptURI
}

// GetLocation returns the value of Location.
func (s *OpCallbackFound) GetLocation() OptURI {
	return s.Location
}

// SetLocation sets the value of Location.
func (s *OpCallbackFound) SetLocation(val OptURI) {
	s.Location = val
}

func (*OpCallbackFound) opCallbackRes() {}

// OpCallbackInternalServerError is response for OpCallback operation.
type OpCallbackInternalServerError struct{}

func (*OpCallbackInternalServerError) opCallbackRes() {}

// OpCertsInternalServerError is response for OpCerts operation.
type OpCertsInternalServerError struct{}

func (*OpCertsInternalServerError) opCertsRes() {}

// OpLoginInternalServerError is response for OpLogin operation.
type OpLoginInternalServerError struct{}

func (*OpLoginInternalServerError) opLoginRes() {}

// OpLoginOK is response for OpLogin operation.
type OpLoginOK struct{}

func (*OpLoginOK) opLoginRes() {}

// OpLoginViewInternalServerError is response for OpLoginView operation.
type OpLoginViewInternalServerError struct{}

func (*OpLoginViewInternalServerError) opLoginViewRes() {}

// OpLoginViewOK is response for OpLoginView operation.
type OpLoginViewOK struct{}

func (*OpLoginViewOK) opLoginViewRes() {}

// OpOpenIDConfigurationInternalServerError is response for OpOpenIDConfiguration operation.
type OpOpenIDConfigurationInternalServerError struct{}

func (*OpOpenIDConfigurationInternalServerError) opOpenIDConfigurationRes() {}

type OpRevokeBadRequest struct {
	// Error.
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *OpRevokeBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *OpRevokeBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*OpRevokeBadRequest) opRevokeRes() {}

// OpRevokeInternalServerError is response for OpRevoke operation.
type OpRevokeInternalServerError struct{}

func (*OpRevokeInternalServerError) opRevokeRes() {}

// OpRevokeOK is response for OpRevoke operation.
type OpRevokeOK struct{}

func (*OpRevokeOK) opRevokeRes() {}

type OpTokenBadRequest struct {
	// Error.
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *OpTokenBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *OpTokenBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*OpTokenBadRequest) opTokenRes() {}

// OpTokenInternalServerError is response for OpToken operation.
type OpTokenInternalServerError struct{}

func (*OpTokenInternalServerError) opTokenRes() {}

// OpUserinfoInternalServerError is response for OpUserinfo operation.
type OpUserinfoInternalServerError struct{}

func (*OpUserinfoInternalServerError) opUserinfoRes() {}

// NewOptIdPSigninRequestSchema returns new OptIdPSigninRequestSchema with value set to v.
func NewOptIdPSigninRequestSchema(v IdPSigninRequestSchema) OptIdPSigninRequestSchema {
	return OptIdPSigninRequestSchema{
		Value: v,
		Set:   true,
	}
}

// OptIdPSigninRequestSchema is optional IdPSigninRequestSchema.
type OptIdPSigninRequestSchema struct {
	Value IdPSigninRequestSchema
	Set   bool
}

// IsSet returns true if OptIdPSigninRequestSchema was set.
func (o OptIdPSigninRequestSchema) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIdPSigninRequestSchema) Reset() {
	var v IdPSigninRequestSchema
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIdPSigninRequestSchema) SetTo(v IdPSigninRequestSchema) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIdPSigninRequestSchema) Get() (v IdPSigninRequestSchema, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIdPSigninRequestSchema) Or(d IdPSigninRequestSchema) IdPSigninRequestSchema {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptIdPSignupRequestSchema returns new OptIdPSignupRequestSchema with value set to v.
func NewOptIdPSignupRequestSchema(v IdPSignupRequestSchema) OptIdPSignupRequestSchema {
	return OptIdPSignupRequestSchema{
		Value: v,
		Set:   true,
	}
}

// OptIdPSignupRequestSchema is optional IdPSignupRequestSchema.
type OptIdPSignupRequestSchema struct {
	Value IdPSignupRequestSchema
	Set   bool
}

// IsSet returns true if OptIdPSignupRequestSchema was set.
func (o OptIdPSignupRequestSchema) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIdPSignupRequestSchema) Reset() {
	var v IdPSignupRequestSchema
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIdPSignupRequestSchema) SetTo(v IdPSignupRequestSchema) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIdPSignupRequestSchema) Get() (v IdPSignupRequestSchema, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIdPSignupRequestSchema) Or(d IdPSignupRequestSchema) IdPSignupRequestSchema {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOPRevokeRequestSchemaTokenTypeHint returns new OptOPRevokeRequestSchemaTokenTypeHint with value set to v.
func NewOptOPRevokeRequestSchemaTokenTypeHint(v OPRevokeRequestSchemaTokenTypeHint) OptOPRevokeRequestSchemaTokenTypeHint {
	return OptOPRevokeRequestSchemaTokenTypeHint{
		Value: v,
		Set:   true,
	}
}

// OptOPRevokeRequestSchemaTokenTypeHint is optional OPRevokeRequestSchemaTokenTypeHint.
type OptOPRevokeRequestSchemaTokenTypeHint struct {
	Value OPRevokeRequestSchemaTokenTypeHint
	Set   bool
}

// IsSet returns true if OptOPRevokeRequestSchemaTokenTypeHint was set.
func (o OptOPRevokeRequestSchemaTokenTypeHint) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOPRevokeRequestSchemaTokenTypeHint) Reset() {
	var v OPRevokeRequestSchemaTokenTypeHint
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOPRevokeRequestSchemaTokenTypeHint) SetTo(v OPRevokeRequestSchemaTokenTypeHint) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOPRevokeRequestSchemaTokenTypeHint) Get() (v OPRevokeRequestSchemaTokenTypeHint, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOPRevokeRequestSchemaTokenTypeHint) Or(d OPRevokeRequestSchemaTokenTypeHint) OPRevokeRequestSchemaTokenTypeHint {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptURI returns new OptURI with value set to v.
func NewOptURI(v url.URL) OptURI {
	return OptURI{
		Value: v,
		Set:   true,
	}
}

// OptURI is optional url.URL.
type OptURI struct {
	Value url.URL
	Set   bool
}

// IsSet returns true if OptURI was set.
func (o OptURI) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptURI) Reset() {
	var v url.URL
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptURI) SetTo(v url.URL) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptURI) Get() (v url.URL, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptURI) Or(d url.URL) url.URL {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// RpCallbackInternalServerError is response for RpCallback operation.
type RpCallbackInternalServerError struct{}

func (*RpCallbackInternalServerError) rpCallbackRes() {}

// RpCallbackOK is response for RpCallback operation.
type RpCallbackOK struct{}

func (*RpCallbackOK) rpCallbackRes() {}

// RpLoginFound is response for RpLogin operation.
type RpLoginFound struct {
	Location OptURI
}

// GetLocation returns the value of Location.
func (s *RpLoginFound) GetLocation() OptURI {
	return s.Location
}

// SetLocation sets the value of Location.
func (s *RpLoginFound) SetLocation(val OptURI) {
	s.Location = val
}

func (*RpLoginFound) rpLoginRes() {}

// RpLoginInternalServerError is response for RpLogin operation.
type RpLoginInternalServerError struct{}

func (*RpLoginInternalServerError) rpLoginRes() {}
