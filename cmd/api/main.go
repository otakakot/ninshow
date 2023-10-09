package main

import (
	"github.com/otakakot/ninshow/internal/adapter/controller"
	"github.com/otakakot/ninshow/internal/adapter/gateway"
	"github.com/otakakot/ninshow/internal/application/interactor"
	"github.com/otakakot/ninshow/internal/driver/config"
	"github.com/otakakot/ninshow/internal/driver/middleware"
	"github.com/otakakot/ninshow/internal/driver/server"
	"github.com/otakakot/ninshow/pkg/api"
)

func main() {
	cfg := config.NewConfig()

	accountRepo := gateway.NewAcccount()

	paramCache := gateway.NewParamCache()

	loggedinCache := gateway.NewLoggedInCache()

	idp := interactor.NewIdentityProvider(accountRepo)

	op := interactor.NewOpenIDProvider(paramCache, loggedinCache, accountRepo)

	rp := interactor.NewRelyingParty()

	ctr := controller.NewController(cfg, idp, op, rp)

	hdl, err := api.NewServer(
		ctr,
		nil,
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
