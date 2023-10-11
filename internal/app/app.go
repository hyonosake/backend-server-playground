package app

import (
	"fmt"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/hyonosake/backend-server-playground/internal/rest_handlers"
	"github.com/hyonosake/backend-server-playground/internal/rest_handlers/default_handler"
)

// TODO: should be stored in .env
const host = "127.0.0.1"
const port = 12321

type App struct {
	Logger          *zap.Logger
	RestRegistrator restRegistrator
}

func NewApp() (App, error) {
	// TODO: .env, configure start params, etc

	logger, err := zap.NewProduction()
	if err != nil {
		return App{}, fmt.Errorf("unable to init logger: %v", err)
	}

	restRegistrar := rest_handlers.NewRestRegistrar()
	restRegistrar.Register(default_handler.DefaultRequestRoute, default_handler.NewDefaultHandler(logger))

	return App{
		Logger:          logger,
		RestRegistrator: restRegistrar,
	}, nil
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	for path, fn := range a.RestRegistrator.List() {
		mux.HandleFunc(path, fn)
	}

	return http.ListenAndServe(":"+strconv.FormatInt(port, 10), mux)
}
