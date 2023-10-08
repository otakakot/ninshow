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

	repo := gateway.NewAcccount()

	uc := interactor.NewAcccount(repo)

	ctr := controller.NewController(uc)

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
