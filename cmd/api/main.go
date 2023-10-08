package main

import (
	"os"

	"github.com/otakakot/ninshow/internal/adapter/controller"
	"github.com/otakakot/ninshow/internal/adapter/gateway"
	"github.com/otakakot/ninshow/internal/application/interactor"
	"github.com/otakakot/ninshow/internal/driver/middleware"
	"github.com/otakakot/ninshow/internal/driver/server"
	"github.com/otakakot/ninshow/pkg/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	accountRepo := gateway.NewAcccount()

	kvs := gateway.NewKVS[any]()

	idp := interactor.NewIdentityProvider(accountRepo)

	op := interactor.NewOpenIDProvider(kvs, accountRepo)

	rp := interactor.NewRelyingParty()

	ctr := controller.NewController(idp, op, rp)

	hdl, err := api.NewServer(
		ctr,
		nil,
	)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(
		port,
		middleware.CORS(hdl),
	)

	srv.Run()
}
