package model

import "golang.org/x/crypto/bcrypt"

type OIDCClient struct {
	ID          string
	Name        string
	Secret      string
	HashSec     []byte
	RedirectURI string
}

func (oc *OIDCClient) CompareSecret(
	secret string,
) error {
	return bcrypt.CompareHashAndPassword(oc.HashSec, []byte(secret))
}

// Deprecated: use local test data.
func GenerateTestOIDCClient(
	id string,
	name string,
	secret string,
	redirectURI string,
) OIDCClient {
	hash, _ := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)

	return OIDCClient{
		ID:          id,
		Name:        name,
		Secret:      secret,
		HashSec:     hash,
		RedirectURI: redirectURI,
	}
}
