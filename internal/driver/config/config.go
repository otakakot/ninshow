package config

import (
	"fmt"
	"os"

	"github.com/otakakot/ninshow/internal/adapter/controller"
)

var _ controller.Config = (*Config)(nil)

type Config struct {
	port         string
	selfEndpoint string
	oidcEndpoint string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	self := os.Getenv("SELF_ENDPOINT")
	if self == "" {
		self = fmt.Sprintf("http://localhost:%s", port)
	}

	oidc := os.Getenv("OIDC_ENDPOINT")
	if oidc == "" {
		oidc = fmt.Sprintf("http://localhost:%s/op", port)
	}

	return &Config{
		port:         port,
		selfEndpoint: self,
		oidcEndpoint: oidc,
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

func (cfg *Config) Port() string {
	return cfg.port
}
