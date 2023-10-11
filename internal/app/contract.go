package app

import "github.com/hyonosake/backend-server-playground/internal/rest_handlers"

type restRegistrator interface {
	Register(route string, handler rest_handlers.Handler)
	List() map[string]rest_handlers.Handler
}
