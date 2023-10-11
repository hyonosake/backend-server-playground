package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"

	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

const host = "127.0.0.1"
const port = 12321

const simpleOrderRequestRoute = "/request/simple_order"
const healthcheckRequestRoute = "/internal/healthcheck"

var orderCount int64

type SimpleOrderRequest struct {
	ProductId int64 `json:"productId"`
}

type Error struct {
	Message string `json:"message"`
}

type Result struct {
	Message string `json:"message"`
	OrderId int64  `json:"orderId"`
}

type SimpleOrderResponse struct {
	Error  *Error  `json:"error,omitempty"`
	Result *Result `json:"result,omitempty"`
}

type HealthcheckResponse struct {
	Alive bool `json:"alive"`
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("unable to init logger: %v", err)
	}
	logger.Sync()

	orderCount++
	http.HandleFunc(simpleOrderRequestRoute, handleSimpleOrderRequest(logger))
	http.HandleFunc(healthcheckRequestRoute, handleHealthcheck())

	strPort := strconv.FormatInt(port, 10)

	err = http.ListenAndServe(":"+strPort, nil)
	if err != nil {
		logger.Fatal("not able to start server",
			zap.Int64("port", port),
			zap.Error(err),
		)
	}
}

func handleSimpleOrderRequest(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SimpleOrderRequest
		var resp SimpleOrderResponse

		w.Header().Set("Content-Type", "application/json")

		defer func() {
			err := json.NewEncoder(w).Encode(resp)
			if err != nil {
				logger.Error("cannot write response", zap.Error(err))
			}
		}()

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			resp.Error = &Error{Message: fmt.Sprintf("method %s not allowed", r.Method)}
			logger.Error("method not allowed", zap.String("method", r.Method))
			return
		}
		body := r.Body
		defer body.Close()
		rawData, err := io.ReadAll(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp.Error = &Error{Message: fmt.Sprintf("method %s not allowed", r.Method)}
			logger.Error("unable to get request data", zap.Error(err))
			return
		}

		err = json.Unmarshal(rawData, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp.Error = &Error{Message: fmt.Sprintf("method %s not allowed", r.Method)}
			logger.Error("invalid request data provided", zap.Error(err))
			return
		}

		resp.Result = &Result{
			Message: fmt.Sprintf("Ммммм, позиция %d для заказа - отличный выбор!", req.ProductId),
			OrderId: atomic.AddInt64(&orderCount, 1),
		}
		w.WriteHeader(http.StatusOK)
	}
}

func handleHealthcheck() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := HealthcheckResponse{Alive: true}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}
