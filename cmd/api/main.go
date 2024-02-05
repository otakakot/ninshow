package main

import (
	"context"
	"fmt"

	"github.com/otakakot/ninshow/internal/adapter/controller"
	"github.com/otakakot/ninshow/internal/adapter/gateway"
	"github.com/otakakot/ninshow/internal/application/interactor"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/driver/config"
	"github.com/otakakot/ninshow/internal/driver/middleware"
	"github.com/otakakot/ninshow/internal/driver/postgres"
	"github.com/otakakot/ninshow/internal/driver/server"
	"github.com/otakakot/ninshow/pkg/api"
)

func main() {
	cfg := config.NewConfig()

	ps := postgres.New()
	defer func() {
		if err := ps.Close(); err != nil {
			panic(err)
		}
	}()

	rdb, err := ps.Of(cfg.DSN)
	if err != nil {
		panic(err)
	}

	accountRepo := gateway.NewAcccount(rdb)

	oidcCliRepo := gateway.NewOIDCClient(rdb)

	// for ReryintParty
	_ = oidcCliRepo.Save(context.Background(), model.GenerateTestOIDCClient( //nolint:staticcheck
		cfg.RelyingPartyID(),
		"ninshow",
		cfg.RelyingPartySecret(),
		fmt.Sprintf("%s/rp/callback", cfg.SelfEndpoint()),
	))

	// for NextAuth
	_ = oidcCliRepo.Save(context.Background(), model.GenerateTestOIDCClient( //nolint:staticcheck
		"f398d470-ce98-4aa9-a39c-fd25c04aa282",
		"ninshow",
		"ninshow",
		"http://localhost:3000/api/auth/callback/ninshow",
	))

	signKeyRepo := gateway.NewJWTSignKey(rdb)

	acc, _ := model.SingupAccount("test", "test@example.com", "test")
	_ = accountRepo.Save(context.Background(), *acc)

	paramCache := gateway.NewParamCache()

	loggedinCache := gateway.NewLoggedInCache()

	atCache := gateway.NewAccessTokenCache()

	rtCache := gateway.NewRefreshTokenCache()

	idp := interactor.NewIdentityProvider(accountRepo)

	op := interactor.NewOpenIDProvider(
		oidcCliRepo,
		accountRepo,
		signKeyRepo,
		paramCache,
		loggedinCache,
		atCache,
		rtCache,
	)

	rp := interactor.NewRelyingParty()

	ctr := controller.NewController(cfg, idp, op, rp)

	secUC := interactor.NewSecurity(atCache)

	sec := controller.NewSecurity(cfg, secUC)

	hdl, err := api.NewServer(
		ctr,
		sec,
		api.WithMiddleware(middleware.Logging()),
	)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(
		cfg.Port(),
		middleware.CORS(hdl),
	)

	srv.Run()
}
