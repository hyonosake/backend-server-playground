package simple_order_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/hyonosake/backend-server-playground/internal/rest_handlers/simple_order_handler/entity"
)

const SimpleOrderRequestRoute = "/orders/simple_order"

var orderCount int64

func Handle(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req entity.SimpleOrderRequest
		var resp entity.SimpleOrderResponse

		w.Header().Set("Content-Type", "application/json")

		defer func() {
			err := json.NewEncoder(w).Encode(resp)
			if err != nil {
				logger.Error("cannot write response", zap.Error(err))
			}
		}()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			resp.Error = &entity.Error{Message: fmt.Sprintf("method %s not allowed", r.Method)}
			logger.Error("method not allowed", zap.String("method", r.Method))
			return
		}

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			resp.Error = &entity.Error{Message: fmt.Sprintf("content type must be application/json, '%s' provided", r.Header.Get("Content-type"))}
			logger.Error("method not allowed", zap.String("method", r.Header.Get("Content-type")))
			return
		}
		body := r.Body
		defer body.Close()
		rawData, err := io.ReadAll(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp.Error = &entity.Error{Message: fmt.Sprintf("method %s not allowed", r.Method)}
			logger.Error("unable to get request data", zap.Error(err))
			return
		}

		err = json.Unmarshal(rawData, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp.Error = &entity.Error{Message: fmt.Sprintf("method %s not allowed", r.Method)}
			logger.Error("invalid request data provided", zap.Error(err))
			return
		}

		resp.Result = &entity.Result{
			Message: fmt.Sprintf("Ммммм, позиция %d для заказа - отличный выбор!", req.ProductId),
			OrderId: atomic.AddInt64(&orderCount, 1),
		}
		w.WriteHeader(http.StatusOK)
	}
}

func NewSimpleOrderHandler(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return Handle(logger)
}
