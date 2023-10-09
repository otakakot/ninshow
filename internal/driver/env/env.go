package env

import (
	"fmt"
	"log/slog"
	"os"
)

type Env string

const (
	prd   Env = "prd"
	stg   Env = "stg"
	dev   Env = "dev"
	local Env = "local"
	empty Env = ""
)

var env Env //nolint:gochecknoglobals

func Init() {
	e := Env(os.Getenv("ENV"))

	slog.Info(fmt.Sprintf("environment: %s", e))

	env = e
}

func Get() Env {
	return env
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsPrd() bool {
	return e == prd
}
