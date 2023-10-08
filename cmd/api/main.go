package main

import (
	"os"

	"github.com/rs/cors"

	"github.com/otakakot/ninshow/internal/adapter/controller"
	"github.com/otakakot/ninshow/internal/driver/server"
	"github.com/otakakot/ninshow/pkg/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctr := &controller.Controller{}

	hdl, err := api.NewServer(
		ctr,
		nil,
	)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(
		port,
		cors.AllowAll().Handler(hdl),
	)

	srv.Run()
}
