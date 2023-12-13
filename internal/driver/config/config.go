package config

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/otakakot/ninshow/internal/adapter/controller"
)

var _ controller.Config = (*Config)(nil)

type Config struct {
	DSN                string
	port               string
	selfEndpoint       string
	oidcEndpoint       string
	idTokenSignKey     *rsa.PrivateKey
	accessTokenSign    string
	relyingPartyID     string
	relyingPartySecret string
}

func NewConfig() *Config {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	self := os.Getenv("SELF_ENDPOINT")
	if self == "" {
		self = fmt.Sprintf("http://localhost:%s", port)
	}

	oidc := os.Getenv("OIDC_ENDPOINT")
	if oidc == "" {
		oidc = fmt.Sprintf("http://localhost:%s/op", port)
	}

	reader := rand.Reader

	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic(err)
	}

	return &Config{
		DSN:                dsn,
		port:               port,
		selfEndpoint:       self,
		oidcEndpoint:       oidc,
		idTokenSignKey:     key,
		accessTokenSign:    "sign",
		relyingPartyID:     "26bf8924-c1d9-484d-8a72-db1df2b05ccd",
		relyingPartySecret: "ninshow",
	}
}

// SelfEndpoint implements controller.Config.
func (cfg Config) SelfEndpoint() string {
	return cfg.selfEndpoint
}

// OIDCEndpoint implements controller.Config.
func (cfg *Config) OIDCEndpoint() string {
	return cfg.oidcEndpoint
}

// IDTokenSignKey implements controller.Config.
func (cfg *Config) IDTokenSignKey() *rsa.PrivateKey {
	return cfg.idTokenSignKey
}

// AcessTokenSign implements controller.Config.
func (cfg *Config) AcessTokenSign() string {
	return cfg.accessTokenSign
}

// RelyingPartyID implements controller.Config.
func (cfg *Config) RelyingPartyID() string {
	return cfg.relyingPartyID
}

// RelyingPartySecret implements controller.Config.
func (cfg *Config) RelyingPartySecret() string {
	return cfg.relyingPartySecret
}

func (cfg *Config) Port() string {
	return cfg.port
}
