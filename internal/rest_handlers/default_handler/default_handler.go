package default_handler

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

const DefaultRequestRoute = "/"

func Handle(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf("Requested resource not found: %s", r.URL.Path)))
		logger.Info("access to undefined resource", zap.String("route", r.URL.Path))
	}
}

func NewDefaultHandler(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return Handle(logger)
}
