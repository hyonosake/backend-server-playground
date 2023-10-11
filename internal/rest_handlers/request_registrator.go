package rest_handlers

import (
	"net/http"

	"github.com/hyonosake/backend-server-playground/internal/utils"
)

type Handler func(http.ResponseWriter, *http.Request)

type restRegistrar struct {
	Handlers map[string]Handler
}

func NewRestRegistrar() *restRegistrar {
	return &restRegistrar{
		Handlers: make(map[string]Handler),
	}
}

func (r *restRegistrar) Register(route string, handler Handler) {
	r.Handlers[route] = handler
}

func (r *restRegistrar) List() map[string]Handler {
	return r.Handlers
}

func (r *restRegistrar) ListRoutes() []string {
	return utils.MapKeys(r.Handlers)
}
