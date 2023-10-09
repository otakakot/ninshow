package main

import (
	"context"

	"github.com/otakakot/ninshow/internal/adapter/controller"
	"github.com/otakakot/ninshow/internal/adapter/gateway"
	"github.com/otakakot/ninshow/internal/application/interactor"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/driver/config"
	"github.com/otakakot/ninshow/internal/driver/middleware"
	"github.com/otakakot/ninshow/internal/driver/server"
	"github.com/otakakot/ninshow/pkg/api"
)

func main() {
	cfg := config.NewConfig()

	accountRepo := gateway.NewAcccount()

	acc, _ := model.SingupAccount("test", "test", "test")
	_ = accountRepo.Save(context.Background(), *acc)

	paramCache := gateway.NewParamCache()

	loggedinCache := gateway.NewLoggedInCache()

	atCache := gateway.NewAccessTokenCache()

	rtCache := gateway.NewRefreshTokenCache()

	idp := interactor.NewIdentityProvider(accountRepo)

	op := interactor.NewOpenIDProvider(
		accountRepo,
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
