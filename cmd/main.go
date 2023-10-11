package main

import (
	"log"

	"github.com/hyonosake/backend-server-playground/internal/app"
	"github.com/hyonosake/backend-server-playground/internal/rest_handlers/healthcheck"
	"github.com/hyonosake/backend-server-playground/internal/rest_handlers/simple_order_handler"
)

func main() {
	app, err := app.NewApp()

	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	app.RestRegistrator.Register(healthcheck.HealthcheckRequestRoute, healthcheck.NewHealthcheckHandler())
	app.RestRegistrator.Register(simple_order_handler.SimpleOrderRequestRoute, simple_order_handler.NewSimpleOrderHandler(app.Logger))

	err = app.Run()

	if err != nil {
		// TODO: graceful
		log.Fatalf("server shutdown: %v", err)
	}
}
